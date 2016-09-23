package testresultparser

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// MVNFailure captures failure noted in Test Case in XML Report
type MVNFailure struct {
	FailureType string `xml:"type,attr"`
	Message     string `xml:",chardata"`
}

// TestCase records test case info in XML Report
type TestCase struct {
	Name string     `xml:"name,attr"`
	F    MVNFailure `xml:"failure"` // struct for Failure Details
}

// TestResultXML is struct for XML Test Report
type TestResultXML struct {
	TestSuite string     `xml:"name,attr"`
	Tests     int        `xml:"tests,attr"`
	Failures  int        `xml:"failures,attr"`
	Errors    int        `xml:"errors,attr"`
	Skipped   int        `xml:"skipped,attr"`
	Time      string     `xml:"time,attr"`
	T         []TestCase `xml:"testcase"` // struct for each Test Case
}

// failsafeReportXML is struct for Failafe Reports
type failsafeReportXML struct {
	Result         string `xml:"result,attr"`
	TimeOut        string `xml:"timeout,attr"`
	Tests          int    `xml:"completed"`
	Failures       int    `xml:"failures"`
	Errors         int    `xml:"errors"`
	Skipped        int    `xml:"skipped"`
	FailureMessage string `xml:"failureMessage"`
}

// Failure struct is the part of the Result struct for the output of the Parser
type Failure struct {
	TestCase string
	Type     string
	Message  string
}

// Result is the output strcut of the Parser
type Result struct {
	TestSuite string
	Tests     int
	Failures  int
	Errors    int
	Skipped   int
	Time      string
	Failure   []Failure
}

// convertResultFromFailsafe is a helper function
// to convert struct to desired output struct
func convertResultFromFailsafe(f failsafeReportXML) Result {
	result := Result{
		TestSuite: f.Result,
		Tests:     f.Tests,
		Failures:  f.Failures,
		Errors:    f.Errors,
		Skipped:   f.Skipped,
		Time:      f.TimeOut,
	}
	result.Failure = append(result.Failure,
		Failure{TestCase: "",
			Type:    "",
			Message: f.FailureMessage,
		})

	return result
}

// convertResultFromXML is a helper function
// to convert struct to desired output struct
func convertResultFromXML(t TestResultXML) Result {
	result := Result{
		TestSuite: t.TestSuite,
		Tests:     t.Tests,
		Failures:  t.Failures,
		Errors:    t.Errors,
		Skipped:   t.Skipped,
		Time:      t.Time}

	for _, test := range t.T {
		if test.F.FailureType != "" {
			result.Failure = append(result.Failure,
				Failure{TestCase: test.Name,
					Type:    test.F.FailureType,
					Message: test.F.Message,
				})
		}
	}
	return result
}

// Parse method for Text Report parsing
func (r *Result) Parse(f []byte) *Result {
	result := Result{}
	failure := Failure{}

	testsuite := regexp.MustCompile("Test set: ([a-zA-Z0-9_.]+)")
	result.TestSuite = testsuite.FindStringSubmatch(string(f))[1]
	tests := regexp.MustCompile("Tests run: ([a-zA-Z0-9_.]+)")
	result.Tests, _ = strconv.Atoi(tests.FindStringSubmatch(string(f))[1])
	failures := regexp.MustCompile("Failures: ([a-zA-Z0-9_.]+)")
	result.Failures, _ = strconv.Atoi(failures.FindStringSubmatch(string(f))[1])
	errors := regexp.MustCompile("Errors: ([a-zA-Z0-9_.]+)")
	result.Errors, _ = strconv.Atoi(errors.FindStringSubmatch(string(f))[1])
	skipped := regexp.MustCompile("Skipped: ([a-zA-Z0-9_.]+)")
	result.Skipped, _ = strconv.Atoi(skipped.FindStringSubmatch(string(f))[1])
	time := regexp.MustCompile("Time elapsed: ([a-zA-Z0-9_.]+)")
	result.Time = time.FindStringSubmatch(string(f))[1]

	if result.Failures != 0 {
		for i := 0; i < result.Failures; i++ {
			failedTestCase := regexp.MustCompile(`.*\(.*\)`)
			failure.TestCase = failedTestCase.FindString(string(f))
			failureType := regexp.MustCompile(`(.*Error)(:.*)`)
			failure.Type = failureType.FindStringSubmatch(string(f))[1]
			failure.Message = failureType.FindStringSubmatch(string(f))[2]
			result.Failure = append(result.Failure, failure)
		}
	}
	return &result
}

// Parse method for XML Report parsing
func (tr *TestResultXML) Parse(b []byte) *Result {
	xml.Unmarshal(b, tr)
	result := convertResultFromXML(*tr)
	return &result
}

// Parse method for Failsafe XML Report parsing
func (fr *failsafeReportXML) Parse(b []byte) *Result {
	xml.Unmarshal(b, fr)
	result := convertResultFromFailsafe(*fr)
	return &result
}

// Parser interface to buildtools parse methods
type Parser interface {
	Parse([]byte) *Result
}

// ReportParser returns byte object to the interface parse method
func ReportParser(p Parser, filepath string) *Result {
	fp, _ := os.Open(filepath)
	defer fp.Close()
	b, _ := ioutil.ReadAll(fp)
	r := p.Parse(b)
	return r
}

// Parse function is called by the external interface and
// returns the output struct
func Parse(filepath, plugin string) (*Result, error) {
	var obj Parser
	if strings.Contains(filepath, ".xml") {
		if plugin == "failsafe" {
			obj = &failsafeReportXML{}
		} else {
			obj = &TestResultXML{}
		}
	} else if strings.Contains(filepath, ".txt") {
		obj = &Result{}
	}
	return ReportParser(obj, filepath), nil
}
