package testresultparser

import (
	"io/ioutil"
	"os"
)

// ReportParser returns byte object to the interface parse method
func readFile(filepath string) []byte {
	fp, _ := os.Open(filepath)
	defer fp.Close()
	b, _ := ioutil.ReadAll(fp)
	return b
}

// convertTestResultFromFailsafe is a helper function
// to convert struct to desired output struct
func convertTestResultFromFailsafe(f failsafeReportXML) TestResult {
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

// convertTestResultFromXML is a helper function
// to convert struct to desired output struct
func convertTestResultFromXML(t testResultXML) TestResult {
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
