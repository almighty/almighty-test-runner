package testresultparser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	assert := assert.New(t)

	trp2 := New("failsafe-summary.xml", "failsafe")
	actualxmlresult1 := trp2.Parse()
	f2 := []Failure{{TestCase: "", Type: "", Message: "Assertion Error"}}
	expectedxmlresult1 := &Result{TestSuite: "254", Tests: 0, Failures: 1, Errors: 0, Skipped: 0, Time: "false", Failure: f2}
	assert.Equal(actualxmlresult1, expectedxmlresult1, "this is expected result")

	trp1 := New("failure_test.xml", "surefire")
	actualxmlresult := trp1.Parse()
	f1 := []Failure{{TestCase: "testApp", Type: "junit.framework.AssertionFailedError", Message: "junit.framework.AssertionFailedError\n\tat junit.framework.Assert.fail(Assert.java:47)\n\tat junit.framework.Assert.assertTrue(Assert.java:20)\n\tat junit.framework.Assert.assertTrue(Assert.java:27)\n\tat org.hemani.javabrains.AppTest.testApp(AppTest.java:36)\n"}}
	expectedxmlresult := &Result{TestSuite: "org.hemani.javabrains.AppTest", Tests: 1, Failures: 1, Errors: 0, Skipped: 0, Time: "0.017", Failure: f1}
	assert.Equal(actualxmlresult, expectedxmlresult, "this is expected result")

	filepath := "failure_test.txt"
	plugin := "surefire"
	trp := New(filepath, plugin)
	actualtxtresult := trp.Parse()
	f := []Failure{{TestCase: "foo(testNgMavenExample1.TestNgMavenExampleTest)", Type: "java.lang.AssertionError", Message: ": expected [foo] but found [1.0]"}}
	expectedtxtresult := &Result{TestSuite: "testNgMavenExample1.TestNgMavenExampleTest", Tests: 2, Failures: 1, Errors: 0, Skipped: 0, Time: "0.785", Failure: f}
	assert.Equal(actualtxtresult, expectedtxtresult, "this is the expected result")
}
