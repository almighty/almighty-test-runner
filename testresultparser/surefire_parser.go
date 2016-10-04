package testresultparser

import "encoding/xml"

// surefireTestFailure captures failure details for the
// corresponding Test Case in XML Report
type surefireTestFailure struct {
	FailureType string `xml:"type,attr"`
	Message     string `xml:",chardata"`
}

// surefireTestCase records information for each of the failing test case in XML Report
type surefireTestCase struct {
	Name string              `xml:"name,attr"`
	F    surefireTestFailure `xml:"failure"` // struct for Failure Details
}

// testResultXML is created as a result of parsing surefire XML Test Report
type surefireReportXML struct {
	TestSuite string             `xml:"name,attr"`
	Tests     int                `xml:"tests,attr"`
	Failures  int                `xml:"failures,attr"`
	Errors    int                `xml:"errors,attr"`
	Skipped   int                `xml:"skipped,attr"`
	Time      string             `xml:"time,attr"`
	T         []surefireTestCase `xml:"testcase"` // struct for each Test Case
}

type surefireParser struct {
}

// convertTestResultFromXML is a helper function to convert
// struct to desired output struct
func (surefireParser) convertTestResultFromXML(t surefireReportXML) TestResult {
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

// Parse method does the actual Surefire XML Report parsing
func (s surefireParser) Parse(filepath string) *TestResult {
	t := surefireReportXML{}
	b := readFile(filepath)
	err := xml.Unmarshal(b, &t)
	checkErr(err)
	TestResult := s.convertTestResultFromXML(t)
	return &TestResult
}
