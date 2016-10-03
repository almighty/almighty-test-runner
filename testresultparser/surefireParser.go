package testresultparser

import "encoding/xml"

// mavenFailure captures failure noted in Test Case in XML Report
type mavenFailure struct {
	FailureType string `xml:"type,attr"`
	Message     string `xml:",chardata"`
}

// TestCase records test case info in XML Report
type testCase struct {
	Name string       `xml:"name,attr"`
	F    mavenFailure `xml:"failure"` // struct for Failure Details
}

// testResultXML is struct for XML Test Report
type testResultXML struct {
	TestSuite string     `xml:"name,attr"`
	Tests     int        `xml:"tests,attr"`
	Failures  int        `xml:"failures,attr"`
	Errors    int        `xml:"errors,attr"`
	Skipped   int        `xml:"skipped,attr"`
	Time      string     `xml:"time,attr"`
	T         []testCase `xml:"testcase"` // struct for each Test Case
}

// Parse method for XML Report parsing
func (tr *testResultXML) Parse(filepath string) *TestResult {
	b := readFile(filepath)
	xml.Unmarshal(b, tr)
	TestResult := convertTestResultFromXML(*tr)
	return &TestResult
}
