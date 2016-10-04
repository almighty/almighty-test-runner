package testresultparser

import (
	"encoding/xml"
	"fmt"
	"os"
)

// gradleTestFailure captures failure details for the
// corresponding Test Case in XML Report
type gradleTestFailure struct {
	FailureType string `xml:"type,attr"`
	Message     string `xml:",chardata"`
}

// gradleTestCase records information for each of the failing test case in XML Report
type gradleTestCase struct {
	Name string            `xml:"name,attr"`
	F    gradleTestFailure `xml:"failure"` // struct for Failure Details
}

// gradleReportXML is created as a result of parsing gradle XML Test Report
type gradleReportXML struct {
	TestSuite string           `xml:"name,attr"`
	Tests     int              `xml:"tests,attr"`
	Failures  int              `xml:"failures,attr"`
	Errors    int              `xml:"errors,attr"`
	Skipped   int              `xml:"skipped,attr"`
	Time      string           `xml:"time,attr"`
	T         []gradleTestCase `xml:"testcase"` // struct for each Test Case
}

type gradleParser struct {
}

// convertTestResultFromXML is a helper function to convert
// struct to desired output struct
func (gradleParser) convertTestResultFromXML(t gradleReportXML) TestResult {
	TestResult := TestResult{
		TestSuite: t.TestSuite,
		Tests:     t.Tests,
		Failures:  t.Failures,
		Errors:    t.Errors,
		Skipped:   t.Skipped,
		Time:      t.Time}

	for _, test := range t.T {
		if test.F.FailureType != "" {
			TestResult.Failure = append(TestResult.Failure,
				TestFailure{TestCase: test.Name,
					Type:    test.F.FailureType,
					Message: test.F.Message,
				})
		}
	}
	return TestResult
}

// Parse method does the actual gradle XML Report parsing
func (s gradleParser) Parse(filepath string) *TestResult {
	t := gradleReportXML{}
	b := readFile(filepath)
	err := xml.Unmarshal(b, &t)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	TestResult := s.convertTestResultFromXML(t)
	return &TestResult
}
