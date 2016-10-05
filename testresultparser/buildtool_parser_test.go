package testresultparser_test

import (
	"fmt"
	"path"
	"path/filepath"
	"runtime"

	. "github.com/almighty/almighty-test-runner/testresultparser"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Parser", func() {

	Describe("Parsing XML file", func() {
		Context("With Maven-surefire plugin and Junit", func() {
			It("should return output struct", func() {
				// given
				filepath := getFilepath("junit_failure_test.xml")
				f := []TestFailure{{TestCase: "testApp", Type: "junit.framework.AssertionFailedError", Message: "junit.framework.AssertionFailedError\n\tat junit.framework.Assert.fail(Assert.java:47)\n\tat junit.framework.Assert.assertTrue(Assert.java:20)\n\tat junit.framework.Assert.assertTrue(Assert.java:27)\n\tat org.hemani.javabrains.AppTest.testApp(AppTest.java:36)\n"}}
				expectedResult := &TestResult{TestSuite: "org.hemani.javabrains.AppTest", Tests: 1, Failures: 1, Errors: 0, Skipped: 0, Time: "0.017", Failure: f}

				// when
				myParser, _ := CreateParser(filepath, "surefire")
				actualResult := myParser.Parse(filepath)

				// then
				Expect(actualResult).To(Equal(expectedResult))
			})
		})
		Context("With Maven-surefire plugin and TestNG", func() {
			It("should return output struct", func() {
				// given
				filepath := getFilepath("testNG_failure_test.xml")
				f := []TestFailure{{TestCase: "foo", Type: "java.lang.AssertionError", Message: "java.lang.AssertionError: expected [foo] but found [1.0]\n\tat org.testng.Assert.fail(Assert.java:94)\n\tat org.testng.Assert.failNotEquals(Assert.java:494)\n\tat org.testng.Assert.assertEquals(Assert.java:123)\n\tat org.testng.Assert.assertEquals(Assert.java:165)\n\tat testNgMavenExample1.TestNgMavenExampleTest.foo(TestNgMavenExampleTest.java:15)\n"}}
				expectedResult := &TestResult{TestSuite: "testNgMavenExample1.TestNgMavenExampleTest", Tests: 2, Failures: 1, Errors: 0, Skipped: 0, Time: "0.009", Failure: f}

				// when
				myParser, _ := CreateParser(filepath, "surefire")
				actualResult := myParser.Parse(filepath)

				// then
				Expect(actualResult).To(Equal(expectedResult))
			})
		})
		Context("With Maven-surefire plugin and TestNG and multiple failures", func() {
			It("should return output struct", func() {
				// given
				filepath := getFilepath("testNG_failure_test2.xml")
				f := []TestFailure{{
					TestCase: "bar",
					Type:     "java.lang.AssertionError",
					Message:  "java.lang.AssertionError: expected [foo] but found [2.0]\n\tat org.testng.Assert.fail(Assert.java:94)\n\tat org.testng.Assert.failNotEquals(Assert.java:494)\n\tat org.testng.Assert.assertEquals(Assert.java:123)\n\tat org.testng.Assert.assertEquals(Assert.java:165)\n\tat testNgMavenExample1.TestNgMavenExampleTest.bar(TestNgMavenExampleTest.java:21)\n",
				},
					{TestCase: "foo",
						Type:    "java.lang.AssertionError",
						Message: "java.lang.AssertionError: expected [foo] but found [1.0]\n\tat org.testng.Assert.fail(Assert.java:94)\n\tat org.testng.Assert.failNotEquals(Assert.java:494)\n\tat org.testng.Assert.assertEquals(Assert.java:123)\n\tat org.testng.Assert.assertEquals(Assert.java:165)\n\tat testNgMavenExample1.TestNgMavenExampleTest.foo(TestNgMavenExampleTest.java:16)\n",
					},
				}
				expectedResult := &TestResult{TestSuite: "testNgMavenExample1.TestNgMavenExampleTest", Tests: 4, Failures: 2, Errors: 0, Skipped: 1, Time: "0.028", Failure: f}

				// when
				myParser, _ := CreateParser(filepath, "surefire")
				actualResult := myParser.Parse(filepath)

				// then
				Expect(actualResult).To(Equal(expectedResult))
			})
		})
		Context("With Maven-failsafe plugin", func() {
			It("should return output struct", func() {
				// given
				filepath := getFilepath("failsafe-summary.xml")
				f := []TestFailure{{TestCase: "", Type: "", Message: "Assertion Error"}}
				expectedResult := &TestResult{TestSuite: "254", Tests: 0, Failures: 1, Errors: 0, Skipped: 0, Time: "false", Failure: f}

				// when
				myParser, _ := CreateParser(filepath, "failsafe")
				actualResult := myParser.Parse(filepath)

				// then
				Expect(actualResult).To(Equal(expectedResult))
			})
		})
		Context("With Gradle default plugin", func() {
			It("should return output struct", func() {
				// given
				filepath := getFilepath("gradle_failure_test.xml")
				f := []TestFailure{{TestCase: "test", Type: "java.lang.AssertionError", Message: "java.lang.AssertionError\n\tat com.udacity.gradle.test.PersonTest.test(PersonTest.java:10)\n\tat sun.reflect.NativeMethodAccessorImpl.invoke0(Native Method)\n\tat sun.reflect.NativeMethodAccessorImpl.invoke(NativeMethodAccessorImpl.java:62)\n\tat sun.reflect.DelegatingMethodAccessorImpl.invoke(DelegatingMethodAccessorImpl.java:43)\n\tat java.lang.reflect.Method.invoke(Method.java:497)\n\tat org.junit.runners.model.FrameworkMethod$1.runReflectiveCall(FrameworkMethod.java:50)\n\tat org.junit.internal.runners.model.ReflectiveCallable.run(ReflectiveCallable.java:12)\n\tat org.junit.runners.model.FrameworkMethod.invokeExplosively(FrameworkMethod.java:47)\n\tat org.junit.internal.runners.statements.InvokeMethod.evaluate(InvokeMethod.java:17)\n\tat org.junit.runners.ParentRunner.runLeaf(ParentRunner.java:325)\n\tat org.junit.runners.BlockJUnit4ClassRunner.runChild(BlockJUnit4ClassRunner.java:78)\n\tat org.junit.runners.BlockJUnit4ClassRunner.runChild(BlockJUnit4ClassRunner.java:57)\n\tat org.junit.runners.ParentRunner$3.run(ParentRunner.java:290)\n\tat org.junit.runners.ParentRunner$1.schedule(ParentRunner.java:71)\n\tat org.junit.runners.ParentRunner.runChildren(ParentRunner.java:288)\n\tat org.junit.runners.ParentRunner.access$000(ParentRunner.java:58)\n\tat org.junit.runners.ParentRunner$2.evaluate(ParentRunner.java:268)\n\tat org.junit.runners.ParentRunner.run(ParentRunner.java:363)\n\tat org.gradle.api.internal.tasks.testing.junit.JUnitTestClassExecuter.runTestClass(JUnitTestClassExecuter.java:112)\n\tat org.gradle.api.internal.tasks.testing.junit.JUnitTestClassExecuter.execute(JUnitTestClassExecuter.java:56)\n\tat org.gradle.api.internal.tasks.testing.junit.JUnitTestClassProcessor.processTestClass(JUnitTestClassProcessor.java:66)\n\tat org.gradle.api.internal.tasks.testing.SuiteTestClassProcessor.processTestClass(SuiteTestClassProcessor.java:51)\n\tat sun.reflect.NativeMethodAccessorImpl.invoke0(Native Method)\n\tat sun.reflect.NativeMethodAccessorImpl.invoke(NativeMethodAccessorImpl.java:62)\n\tat sun.reflect.DelegatingMethodAccessorImpl.invoke(DelegatingMethodAccessorImpl.java:43)\n\tat java.lang.reflect.Method.invoke(Method.java:497)\n\tat org.gradle.messaging.dispatch.ReflectionDispatch.dispatch(ReflectionDispatch.java:35)\n\tat org.gradle.messaging.dispatch.ReflectionDispatch.dispatch(ReflectionDispatch.java:24)\n\tat org.gradle.messaging.dispatch.ContextClassLoaderDispatch.dispatch(ContextClassLoaderDispatch.java:32)\n\tat org.gradle.messaging.dispatch.ProxyDispatchAdapter$DispatchingInvocationHandler.invoke(ProxyDispatchAdapter.java:93)\n\tat com.sun.proxy.$Proxy2.processTestClass(Unknown Source)\n\tat org.gradle.api.internal.tasks.testing.worker.TestWorker.processTestClass(TestWorker.java:109)\n\tat sun.reflect.NativeMethodAccessorImpl.invoke0(Native Method)\n\tat sun.reflect.NativeMethodAccessorImpl.invoke(NativeMethodAccessorImpl.java:62)\n\tat sun.reflect.DelegatingMethodAccessorImpl.invoke(DelegatingMethodAccessorImpl.java:43)\n\tat java.lang.reflect.Method.invoke(Method.java:497)\n\tat org.gradle.messaging.dispatch.ReflectionDispatch.dispatch(ReflectionDispatch.java:35)\n\tat org.gradle.messaging.dispatch.ReflectionDispatch.dispatch(ReflectionDispatch.java:24)\n\tat org.gradle.messaging.remote.internal.hub.MessageHub$Handler.run(MessageHub.java:360)\n\tat org.gradle.internal.concurrent.ExecutorPolicy$CatchAndRecordFailures.onExecute(ExecutorPolicy.java:54)\n\tat org.gradle.internal.concurrent.StoppableExecutorImpl$1.run(StoppableExecutorImpl.java:40)\n\tat java.util.concurrent.ThreadPoolExecutor.runWorker(ThreadPoolExecutor.java:1142)\n\tat java.util.concurrent.ThreadPoolExecutor$Worker.run(ThreadPoolExecutor.java:617)\n\tat java.lang.Thread.run(Thread.java:745)\n"}}
				expectedResult := &TestResult{TestSuite: "com.udacity.gradle.test.PersonTest", Tests: 1, Failures: 1, Errors: 0, Skipped: 0, Time: "0.003", Failure: f}

				// when
				myParser, _ := CreateParser(filepath, "gradle")
				actualResult := myParser.Parse(filepath)

				//then
				Expect(actualResult).To(Equal(expectedResult))
			})
		})

	})
	Describe("Parsing TXT file", func() {
		Context("With Maven-surefire plugin and TestNG", func() {
			It("should return output struct", func() {
				// given
				filepath := getFilepath("testNG_failure_test.txt")
				f := []TestFailure{{TestCase: "foo", Type: "java.lang.AssertionError", Message: "expected [foo] but found [1.0]"}}
				expectedResult := &TestResult{TestSuite: "testNgMavenExample1.TestNgMavenExampleTest", Tests: 2, Failures: 1, Errors: 0, Skipped: 0, Time: "0.785", Failure: f}

				// when
				myParser, _ := CreateParser(filepath, "surefire")
				actualResult := myParser.Parse(filepath)

				//then
				Expect(actualResult).To(Equal(expectedResult))
			})
		})
		Context("With Maven-surefire plugin and TestNG and multiple failures", func() {
			It("should return output struct", func() {
				// given
				filepath := getFilepath("testNG_failure_test2.txt")
				f := []TestFailure{{
					TestCase: "bar",
					Type:     "java.lang.AssertionError",
					Message:  "expected [foo] but found [2.0]",
				},
					{
						TestCase: "foo",
						Type:     "java.lang.AssertionError",
						Message:  "expected [foo] but found [1.0]",
					},
				}
				expectedResult := &TestResult{TestSuite: "testNgMavenExample1.TestNgMavenExampleTest", Tests: 4, Failures: 2, Errors: 0, Skipped: 1, Time: "1.395", Failure: f}

				// when
				myParser, _ := CreateParser(filepath, "surefire")
				actualResult := myParser.Parse(filepath)

				// then
				Expect(actualResult).To(Equal(expectedResult))
			})
		})
	})
})

// --- Test helpers

func getFilepath(filename string) string {
	_, packagefilename, _, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("Can't find filepath.")
	}
	targetDir := filepath.FromSlash(path.Dir(packagefilename) + "/test_fixtures/")
	targetPath := filepath.FromSlash(targetDir + filename)
	return targetPath
}
