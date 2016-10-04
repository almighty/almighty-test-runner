package testresultparser

import "encoding/xml"

// failsafeReportXML is created as result of parsing Failafe Reports
type failsafeReportXML struct {
	Result         string `xml:"result,attr"`
	TimeOut        string `xml:"timeout,attr"`
	Tests          int    `xml:"completed"`
	Failures       int    `xml:"failures"`
	Errors         int    `xml:"errors"`
	Skipped        int    `xml:"skipped"`
	FailureMessage string `xml:"failureMessage"`
}

type failsafeParser struct {
}

// convertTestResultFromFailsafe is a helper function
// to convert struct to desired output struct
func (failsafeParser) convertTestResultFromFailsafe(f failsafeReportXML) TestResult {
	TestResult := TestResult{
		TestSuite: f.Result,
		Tests:     f.Tests,
		Failures:  f.Failures,
		Errors:    f.Errors,
		Skipped:   f.Skipped,
		Time:      f.TimeOut,
	}
	TestResult.Failure = append(TestResult.Failure,
		TestFailure{TestCase: "",
			Type:    "",
			Message: f.FailureMessage,
		})
	return TestResult
}

// Parse method does the actual Failsafe XML Report parsing
func (p failsafeParser) Parse(filepath string) *TestResult {
	f := failsafeReportXML{}
	b := readFile(filepath)
	err := xml.Unmarshal(b, &f)
	checkErr(err)
	TestResult := p.convertTestResultFromFailsafe(f)
	return &TestResult
}
