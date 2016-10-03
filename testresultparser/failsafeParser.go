package testresultparser

import "encoding/xml"

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

// Parse method for Failsafe XML Report parsing
func (fr *failsafeReportXML) Parse(filepath string) *TestResult {
	b := readFile(filepath)
	xml.Unmarshal(b, fr)
	TestResult := convertTestResultFromFailsafe(*fr)
	return &TestResult
}
