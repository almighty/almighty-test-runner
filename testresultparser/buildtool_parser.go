package testresultparser

import (
	"errors"
	"strings"
)

// TestFailure struct is the part of the TestResult struct for the output of the Parser
type TestFailure struct {
	TestCase string
	Type     string
	Message  string
}

// TestResult is the output struct of the Parser
type TestResult struct {
	TestSuite string
	Tests     int
	Failures  int
	Errors    int
	Skipped   int
	Time      string
	Failure   []TestFailure
}

// Parser interface to buildtools parse methods
type Parser interface {
	Parse(filepath string) *TestResult
}

// CreateParser function is called by the external interface and
// returns the output struct
func CreateParser(filepath, plugin string) (Parser, error) {
	switch plugin {
	case "surefire":
		if strings.Contains(filepath, ".xml") {
			return new(surefireParser), nil
		} else if strings.Contains(filepath, ".txt") {
			return new(textParser), nil
		}
	case "failsafe":
		return new(failsafeParser), nil
	case "gradle":
		return new(gradleParser), nil
	}
	return nil, errors.New("Invalid Parser Type.")
}
