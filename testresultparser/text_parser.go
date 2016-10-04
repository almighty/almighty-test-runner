package testresultparser

import (
	"regexp"
	"strconv"
)

type textParser struct {
}

// Parse method parses the Surefire Text Report
func (textParser) Parse(filepath string) *TestResult {
	testResult := TestResult{}

	f := readFile(filepath)
	testsuite := regexp.MustCompile("Test set: ([a-zA-Z0-9_.]+)")
	tests := regexp.MustCompile("Tests run: ([a-zA-Z0-9_.]+)")
	failures := regexp.MustCompile("Failures: ([a-zA-Z0-9_.]+)")
	errors := regexp.MustCompile("Errors: ([a-zA-Z0-9_.]+)")
	skipped := regexp.MustCompile("Skipped: ([a-zA-Z0-9_.]+)")
	time := regexp.MustCompile("Time elapsed: ([a-zA-Z0-9_.]+)")

	testResult.TestSuite = testsuite.FindStringSubmatch(string(f))[1]
	testResult.Tests, _ = strconv.Atoi(tests.FindStringSubmatch(string(f))[1])
	testResult.Failures, _ = strconv.Atoi(failures.FindStringSubmatch(string(f))[1])
	testResult.Errors, _ = strconv.Atoi(errors.FindStringSubmatch(string(f))[1])
	testResult.Skipped, _ = strconv.Atoi(skipped.FindStringSubmatch(string(f))[1])
	testResult.Time = time.FindStringSubmatch(string(f))[1]

	if testResult.Failures != 0 {
		failedTestCase := regexp.MustCompile(`(.*)\(.*\).*FAILURE`)
		failureType := regexp.MustCompile(`(.*Error): (.*)`)
		for i := 0; i < testResult.Failures; i++ {
			testResult.Failure = append(testResult.Failure,
				TestFailure{TestCase: failedTestCase.FindAllStringSubmatch(string(f), -1)[i][1],
					Type:    failureType.FindAllStringSubmatch(string(f), -1)[i][1],
					Message: failureType.FindAllStringSubmatch(string(f), -1)[i][2],
				})
		}
	}
	return &testResult
}
