package testresultparser_test

import (
	. "github.com/almighty-test-runner/testresultparser"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Parser", func() {

	var (
		path           = "/home/hemani/workspace/src/github.com/almighty-test-runner/testresultparser/artifacts/"
		expectedResult *Result
	)
	/*
		BeforeEach(func() {
	})*/

	Describe("Parsing XML file", func() {
		Context("With Maven-surefire plugin and Junit", func() {
			It("should return output struct", func() {
				f := []Failure{{TestCase: "testApp", Type: "junit.framework.AssertionFailedError", Message: "junit.framework.AssertionFailedError\n\tat junit.framework.Assert.fail(Assert.java:47)\n\tat junit.framework.Assert.assertTrue(Assert.java:20)\n\tat junit.framework.Assert.assertTrue(Assert.java:27)\n\tat org.hemani.javabrains.AppTest.testApp(AppTest.java:36)\n"}}
				expectedResult = &Result{TestSuite: "org.hemani.javabrains.AppTest", Tests: 1, Failures: 1, Errors: 0, Skipped: 0, Time: "0.017", Failure: f}
				Expect(Parse(path+"junit_failure_test.xml", "surefire")).To(Equal(expectedResult))
			})
		})
		Context("With Maven-surefire plugin and TestNG", func() {
			It("should return output struct", func() {
				f := []Failure{{TestCase: "foo", Type: "java.lang.AssertionError", Message: "java.lang.AssertionError: expected [foo] but found [1.0]\n\tat org.testng.Assert.fail(Assert.java:94)\n\tat org.testng.Assert.failNotEquals(Assert.java:494)\n\tat org.testng.Assert.assertEquals(Assert.java:123)\n\tat org.testng.Assert.assertEquals(Assert.java:165)\n\tat testNgMavenExample1.TestNgMavenExampleTest.foo(TestNgMavenExampleTest.java:15)\n"}}
				expectedResult = &Result{TestSuite: "testNgMavenExample1.TestNgMavenExampleTest", Tests: 2, Failures: 1, Errors: 0, Skipped: 0, Time: "0.009", Failure: f}
				Expect(Parse(path+"testNG_failure_test.xml", "surefire")).To(Equal(expectedResult))
			})
		})
		Context("With Maven-failsafe plugin", func() {
			It("should return output struct", func() {
				f := []Failure{{TestCase: "", Type: "", Message: "Assertion Error"}}
				expectedResult = &Result{TestSuite: "254", Tests: 0, Failures: 1, Errors: 0, Skipped: 0, Time: "false", Failure: f}
				Expect(Parse(path+"failsafe-summary.xml", "failsafe")).To(Equal(expectedResult))
			})
		})
		Context("With Gradle default plugin", func() {
			It("should return output struct", func() {
				f := []Failure{{TestCase: "test", Type: "java.lang.AssertionError", Message: "java.lang.AssertionError\n\tat com.udacity.gradle.test.PersonTest.test(PersonTest.java:10)\n\tat sun.reflect.NativeMethodAccessorImpl.invoke0(Native Method)\n\tat sun.reflect.NativeMethodAccessorImpl.invoke(NativeMethodAccessorImpl.java:62)\n\tat sun.reflect.DelegatingMethodAccessorImpl.invoke(DelegatingMethodAccessorImpl.java:43)\n\tat java.lang.reflect.Method.invoke(Method.java:497)\n\tat org.junit.runners.model.FrameworkMethod$1.runReflectiveCall(FrameworkMethod.java:50)\n\tat org.junit.internal.runners.model.ReflectiveCallable.run(ReflectiveCallable.java:12)\n\tat org.junit.runners.model.FrameworkMethod.invokeExplosively(FrameworkMethod.java:47)\n\tat org.junit.internal.runners.statements.InvokeMethod.evaluate(InvokeMethod.java:17)\n\tat org.junit.runners.ParentRunner.runLeaf(ParentRunner.java:325)\n\tat org.junit.runners.BlockJUnit4ClassRunner.runChild(BlockJUnit4ClassRunner.java:78)\n\tat org.junit.runners.BlockJUnit4ClassRunner.runChild(BlockJUnit4ClassRunner.java:57)\n\tat org.junit.runners.ParentRunner$3.run(ParentRunner.java:290)\n\tat org.junit.runners.ParentRunner$1.schedule(ParentRunner.java:71)\n\tat org.junit.runners.ParentRunner.runChildren(ParentRunner.java:288)\n\tat org.junit.runners.ParentRunner.access$000(ParentRunner.java:58)\n\tat org.junit.runners.ParentRunner$2.evaluate(ParentRunner.java:268)\n\tat org.junit.runners.ParentRunner.run(ParentRunner.java:363)\n\tat org.gradle.api.internal.tasks.testing.junit.JUnitTestClassExecuter.runTestClass(JUnitTestClassExecuter.java:112)\n\tat org.gradle.api.internal.tasks.testing.junit.JUnitTestClassExecuter.execute(JUnitTestClassExecuter.java:56)\n\tat org.gradle.api.internal.tasks.testing.junit.JUnitTestClassProcessor.processTestClass(JUnitTestClassProcessor.java:66)\n\tat org.gradle.api.internal.tasks.testing.SuiteTestClassProcessor.processTestClass(SuiteTestClassProcessor.java:51)\n\tat sun.reflect.NativeMethodAccessorImpl.invoke0(Native Method)\n\tat sun.reflect.NativeMethodAccessorImpl.invoke(NativeMethodAccessorImpl.java:62)\n\tat sun.reflect.DelegatingMethodAccessorImpl.invoke(DelegatingMethodAccessorImpl.java:43)\n\tat java.lang.reflect.Method.invoke(Method.java:497)\n\tat org.gradle.messaging.dispatch.ReflectionDispatch.dispatch(ReflectionDispatch.java:35)\n\tat org.gradle.messaging.dispatch.ReflectionDispatch.dispatch(ReflectionDispatch.java:24)\n\tat org.gradle.messaging.dispatch.ContextClassLoaderDispatch.dispatch(ContextClassLoaderDispatch.java:32)\n\tat org.gradle.messaging.dispatch.ProxyDispatchAdapter$DispatchingInvocationHandler.invoke(ProxyDispatchAdapter.java:93)\n\tat com.sun.proxy.$Proxy2.processTestClass(Unknown Source)\n\tat org.gradle.api.internal.tasks.testing.worker.TestWorker.processTestClass(TestWorker.java:109)\n\tat sun.reflect.NativeMethodAccessorImpl.invoke0(Native Method)\n\tat sun.reflect.NativeMethodAccessorImpl.invoke(NativeMethodAccessorImpl.java:62)\n\tat sun.reflect.DelegatingMethodAccessorImpl.invoke(DelegatingMethodAccessorImpl.java:43)\n\tat java.lang.reflect.Method.invoke(Method.java:497)\n\tat org.gradle.messaging.dispatch.ReflectionDispatch.dispatch(ReflectionDispatch.java:35)\n\tat org.gradle.messaging.dispatch.ReflectionDispatch.dispatch(ReflectionDispatch.java:24)\n\tat org.gradle.messaging.remote.internal.hub.MessageHub$Handler.run(MessageHub.java:360)\n\tat org.gradle.internal.concurrent.ExecutorPolicy$CatchAndRecordFailures.onExecute(ExecutorPolicy.java:54)\n\tat org.gradle.internal.concurrent.StoppableExecutorImpl$1.run(StoppableExecutorImpl.java:40)\n\tat java.util.concurrent.ThreadPoolExecutor.runWorker(ThreadPoolExecutor.java:1142)\n\tat java.util.concurrent.ThreadPoolExecutor$Worker.run(ThreadPoolExecutor.java:617)\n\tat java.lang.Thread.run(Thread.java:745)\n"}}
				expectedResult = &Result{TestSuite: "com.udacity.gradle.test.PersonTest", Tests: 1, Failures: 1, Errors: 0, Skipped: 0, Time: "0.003", Failure: f}
				Expect(Parse(path+"gradle_failure_test.xml", "default")).To(Equal(expectedResult))
			})
		})

	})
	Describe("Parsing TXT file", func() {
		Context("With Maven-surefire plugin and TestNG", func() {
			It("should return output struct", func() {
				f := []Failure{{TestCase: "foo(testNgMavenExample1.TestNgMavenExampleTest)", Type: "java.lang.AssertionError", Message: ": expected [foo] but found [1.0]"}}
				expectedResult = &Result{TestSuite: "testNgMavenExample1.TestNgMavenExampleTest", Tests: 2, Failures: 1, Errors: 0, Skipped: 0, Time: "0.785", Failure: f}
				Expect(Parse(path+"testNG_failure_test.txt", "surefire")).To(Equal(expectedResult))
			})
		})
	})
})
