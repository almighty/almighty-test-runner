package testresultparser

// TestResultKind represents what is the status of the test execution
type TestResultKind uint8

const (
	PASSED TestResultKind = iota + 1
	SKIPPED
	FAILURE
	ERROR
)

// TestResult contains details of single test execution
type TestResult struct {
	TestCase,
	Time string
	Kind TestResultKind
	Message,
	Type,
	Details string
}

// ExecutionSummary keeps overall information about test suite execution
type ExecutionSummary struct {
	Total,
	Failures,
	Errors,
	Skipped int
	Time string
}

// TestResults contains details of the whole test suite execution including subsequent
type TestResults struct {
	Name    string
	Summary ExecutionSummary
	Results []TestResult
}

// Parser interface provides a way to read report file generated by a test runner
type Parser interface {
	Parse(filepath string) (*TestResults, error)
}
