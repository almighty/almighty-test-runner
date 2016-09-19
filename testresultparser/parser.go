package testresultparser

import (
	"encoding/xml"
	"fmt"
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

// Failure struct is the part of the Result struct for the output of the Parser Library
type Failure struct {
	TestCase string
	Type     string
	Message  string
}

// Result is the output of the Parser Library
type Result struct {
	TestSuite string
	Tests     int
	Failures  int
	Errors    int
	Skipped   int
	Time      string
	Failure   []Failure
}

// Parser struct is the input for the Parser Library
type Parser struct {
	filepath string
	plugin   string
}

func parsefailsafeXML(filename string) (failsafeReportXML, error) {
	v := failsafeReportXML{}
	xmlFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return v, err
	}
	defer xmlFile.Close()
	b, _ := ioutil.ReadAll(xmlFile)
	xml.Unmarshal(b, &v)
	return v, nil
}

func parseXML(filename string) (TestResultXML, error) {
	v := TestResultXML{}
	xmlFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return v, err
	}
	defer xmlFile.Close()
	b, _ := ioutil.ReadAll(xmlFile)
	xml.Unmarshal(b, &v)
	return v, nil
}

func parseTxt(filename string) (Result, error) {
	result := Result{}
	failure := Failure{}
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return result, err
	}

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
	return result, nil
}

func convertResultFromFailsafe(f failsafeReportXML) Result {
	var result Result
	var failure Failure

	result.TestSuite = f.Result
	result.Tests = f.Tests
	result.Failures = f.Failures
	result.Errors = f.Errors
	result.Skipped = f.Skipped
	result.Time = f.TimeOut
	failure.TestCase = ""
	failure.Type = ""
	failure.Message = f.FailureMessage
	result.Failure = append(result.Failure, failure)

	return result
}

func convertResultFromXML(t TestResultXML) Result {
	var result Result
	var failure Failure

	result.TestSuite = t.TestSuite
	result.Tests = t.Tests
	result.Failures = t.Failures
	result.Errors = t.Errors
	result.Skipped = t.Skipped
	result.Time = t.Time
	for _, test := range t.T {
		failure.TestCase = test.Name
		failure.Type = test.F.FailureType
		failure.Message = test.F.Message
	}
	result.Failure = append(result.Failure, failure)
	return result
}

// Parse function defines the parsing function used for the specified format
func (p Parser) Parse() *Result {
	var result Result
	if strings.Contains(p.filepath, ".xml") {
		if p.plugin == "failsafe" {
			r, _ := parsefailsafeXML(p.filepath)
			result = convertResultFromFailsafe(r)
		} else {
			r, _ := parseXML(p.filepath)
			result = convertResultFromXML(r)
		}
	}
	if strings.Contains(p.filepath, ".txt") {
		result, _ = parseTxt(p.filepath)
	}
	return &result
}

// New function acts like a constructor
func New(filepath, plugin string) *Parser {
	p := Parser{filepath: filepath, plugin: plugin}
	return &p
}
