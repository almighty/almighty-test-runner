package testresultparser

import "encoding/xml"

type report struct {
	TestSuite   string     `xml:"name,attr"`
	Tests       int        `xml:"tests,attr"`
	Failures    int        `xml:"failures,attr"`
	Errors      int        `xml:"errors,attr"`
	Skipped     int        `xml:"skipped,attr"`
	Time        float64    `xml:"time,attr"`
	TestResults []testCase `xml:"testcase"`
	StdOut      string     `xml:"system-out",chardata`
	StdErr      string     `xml:"system-err",chardata`
}

// surefireTestCase records information for each of the failing test case in XML Report
type testCase struct {
	Name   string  `xml:"name,attr"`
	Time   float64 `xml:"time,attr"`
	Report []tag   `xml:",any"`
	StdOut string  `xml:"system-out",chardata`
	StdErr string  `xml:"system-err",chardata`
}

type tag struct {
	XMLName xml.Name `xml:""`
	Type    string   `xml:"type,attr"`
	Message string   `xml:"message,attr"`
	Content string   `xml:",chardata"`
}

type extractor func(tag) string

var getType = func(t tag) string { return t.Type }
var getMsg = func(t tag) string { return t.Message }
var getContent = func(t tag) string { return t.Content }

func convertTestResultFromJUnitFormat(t *report) TestResults {
	results := TestResults{
		Name: t.TestSuite,
		Summary: ExecutionSummary{
			Total:     t.Tests,
			Failures:  t.Failures,
			Errors:    t.Errors,
			Skipped:   t.Skipped,
			Time:      t.Time,
			SystemOut: t.StdOut,
			SystemErr: t.StdErr,
		},
	}

	for _, test := range t.TestResults {
		results.Results = append(results.Results,
			TestResult{
				TestCase: test.Name,
				Time:     test.Time,
				Kind:     toResultType(test),
				Type:     detailsOf(test, getType),
				Message:  detailsOf(test, getMsg),
				Details:  detailsOf(test, getContent),
			})
		results.Summary.SystemOut += test.StdOut
		results.Summary.SystemErr += test.StdErr
	}
	return results
}

func detailsOf(t testCase, ext extractor) string {
	var result = ""

	if t.isError() {
		result = extract(t, "error", ext)
	}

	if t.isFailure() {
		result = extract(t, "failure", ext)
	}

	return result
}

func extract(t testCase, tagName string, ext extractor) string {
	for _, element := range t.Report {
		if element.XMLName.Local == tagName {
			return ext(element)
		}
	}
	return ""
}

func toResultType(testCase testCase) TestResultKind {
	if testCase.isSkipped() {
		return SKIPPED
	}

	if testCase.isError() {
		return ERROR
	}

	if testCase.isFailure() {
		return FAILURE
	}

	return PASSED
}

func (s testCase) isSkipped() bool {
	return s.isType("skipped")
}

func (s testCase) isError() bool {
	return s.isType("error")
}

func (s testCase) isFailure() bool {
	return s.isType("failure")
}

func (s testCase) isType(t string) bool {

	for _, element := range s.Report {
		if element.XMLName.Local == t {
			return true
		}
	}
	return false
}
