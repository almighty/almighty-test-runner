package testresultparser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseXML(t *testing.T) {
	assert := assert.New(t)
	trp := New("failure_test.xml", "surefire")
	actualresult := trp.Parse()
	f := []Failure{{TestCase: "testApp", Type: "junit.framework.AssertionFailedError", Message: "junit.framework.AssertionFailedError\n\tat junit.framework.Assert.fail(Assert.java:47)\n\tat junit.framework.Assert.assertTrue(Assert.java:20)\n\tat junit.framework.Assert.assertTrue(Assert.java:27)\n\tat org.hemani.javabrains.AppTest.testApp(AppTest.java:36)\n"}}
	expectedresult := &Result{TestSuite: "org.hemani.javabrains.AppTest", Tests: 1, Failures: 1, Errors: 0, Skipped: 0, Time: "0.017", Failure: f}
	assert.Equal(actualresult, expectedresult, "this is expected result")
}

func TestParseFailsafeXML(t *testing.T) {
	assert := assert.New(t)
	trp := New("failsafe-summary.xml", "failsafe")
	actualresult := trp.Parse()
	f := []Failure{{TestCase: "", Type: "", Message: "Assertion Error"}}
	expectedresult := &Result{TestSuite: "254", Tests: 0, Failures: 1, Errors: 0, Skipped: 0, Time: "false", Failure: f}
	assert.Equal(actualresult, expectedresult, "this is expected result")
}

func TestParseTxt(t *testing.T) {
	assert := assert.New(t)
	filepath := "failure_test.txt"
	plugin := "surefire"
	trp := New(filepath, plugin)
	actualresult := trp.Parse()
	f := []Failure{{TestCase: "foo(testNgMavenExample1.TestNgMavenExampleTest)", Type: "java.lang.AssertionError", Message: ": expected [foo] but found [1.0]"}}
	expectedresult := &Result{TestSuite: "testNgMavenExample1.TestNgMavenExampleTest", Tests: 2, Failures: 1, Errors: 0, Skipped: 0, Time: "0.785", Failure: f}
	assert.Equal(actualresult, expectedresult, "this is the expected result")
}
