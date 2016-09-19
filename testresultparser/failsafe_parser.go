package testresultparser

import "encoding/xml"

// FailsafeParser provides a mechanism to parse Maven Failsafe summary report
type FailsafeParser struct {}

// Parse method does the actual Failsafe XML Report parsing
func (p *FailsafeParser) Parse(filepath string) (*TestResults, error) {
	var f *failsafeReportXML
	b := readFile(filepath)
	if err := xml.Unmarshal(b, &f); err != nil {
		return nil, err
	}
	testResult := convertTestResultFromFailsafeSummary(f)
	return &testResult, nil
}

// failsafeReportXML is created as result of parsing Failsafe Summary Reports
type failsafeReportXML struct {
	Result         string `xml:"result,attr"`
	TimeOut        string `xml:"timeout,attr"`
	Tests          int    `xml:"completed"`
	Failures       int    `xml:"failures"`
	Errors         int    `xml:"errors"`
	Skipped        int    `xml:"skipped"`
	FailureMessage string `xml:"failureMessage"`
}

func convertTestResultFromFailsafeSummary(f *failsafeReportXML) TestResults {
	testResult := TestResults{
		Name: f.Result,
		Summary: ExecutionSummary{
			Total:     f.Tests,
			Failures:  f.Failures,
			Errors:    f.Errors,
			Skipped:   f.Skipped,
		},
	}

	return testResult
}
