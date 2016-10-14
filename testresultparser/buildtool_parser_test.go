package testresultparser_test

import (
	"path"
	"path/filepath"
	"runtime"

	. "github.com/almighty/almighty-test-runner/testresultparser"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Parser", func() {

	Describe("Parsing Maven Surefire XML report", func() {

		It("should have failures in the results", func() {
			// given
			filepath := getFilepath("single_test_failure_testng_surefire.xml")
			expectedFailure := TestResult{
				TestCase: "foo",
				Time:     "0.003",
				Kind:     FAILURE,
				Type:     "java.lang.AssertionError",
				Message:  "expected [foo] but found [1.0]",
				Details:  "java.lang.AssertionError: expected [foo] but found [1.0]\n\tat org.testng.Assert.fail(Assert.java:94)\n\tat org.testng.Assert.failNotEquals(Assert.java:494)\n\tat org.testng.Assert.assertEquals(Assert.java:123)\n\tat org.testng.Assert.assertEquals(Assert.java:165)\n\tat testNgMavenExample1.TestNgMavenExampleTest.foo(TestNgMavenExampleTest.java:15)\n",
			}
			surefireParser := SurefireParser{}

			// when
			actualResult, _ := surefireParser.Parse(filepath)

			// then
			Expect(actualResult.Results).To(ContainElement(expectedFailure))
		})

		It("should parse summary", func() {
			// given
			filepath := getFilepath("single_test_failure_testng_surefire.xml")
			expectedSummary := ExecutionSummary{
				Total: 2, Failures: 1, Errors: 0, Skipped: 0, Time: "0.009", SystemOut: "", SystemErr: "",
			}
			surefireParser := SurefireParser{}

			// when
			actualResult, _ := surefireParser.Parse(filepath)

			// then
			Expect(actualResult.Summary).To(Equal(expectedSummary))
		})

		It("should convert to TestResult when multiple failures reported", func() {
			// given
			filepath := getFilepath("multiple_test_failures_junit_surefire.xml")
			failure := TestResult{
				TestCase: "shouldUseHostEnvIfDockerHostIsSetOnServerURIAndSystemEnvironmentVarIsSet",
				Time:     "0.002",
				Kind:     FAILURE,
				Message:  "\nExpected: map containing [\"serverUri\"->\"tcp://127.0.0.1:22222\"]\n     but: map was [<serverUri=tcp://localhost:4243>, <dockerServerIp=localhost>, <tlsVerify=false>]",
				Type:     "java.lang.AssertionError",
				Details:  "java.lang.AssertionError:\nExpected: map containing [\"serverUri\"->\"tcp://127.0.0.1:22222\"]\n     but: map was [<serverUri=tcp://localhost:4243>, <dockerServerIp=localhost>, <tlsVerify=false>]\n\tat org.hamcrest.MatcherAssert.assertThat(MatcherAssert.java:20)\n\tat org.junit.Assert.assertThat(Assert.java:865)\n\tat org.junit.Assert.assertThat(Assert.java:832)\n\tat org.arquillian.cube.docker.impl.client.CubeConfiguratorTest.shouldUseHostEnvIfDockerHostIsSetOnServerURIAndSystemEnvironmentVarIsSet(CubeConfiguratorTest.java:260)\n",
			}

			error := TestResult{
				TestCase: "shouldSetServerIpWithLocalhostInCaseOfNativeLinuxInstallation",
				Time:     "0.003",
				Kind:     ERROR,
				Type:     "java.lang.NullPointerException:",
				Details:  "java.lang.NullPointerException: null\n      at org.arquillian.cube.docker.impl.client.CubeConfiguratorTest.shouldSetServerIpWithLocalhostInCaseOfNativeLinuxInstallation(CubeConfiguratorTest.java:226)\n    ",
			}

			surefireParser := SurefireParser{}

			// when
			actualResult, _ := surefireParser.Parse(filepath)

			// then
			Expect(actualResult.Results).To(HaveLen(22))
			Expect(actualResult.Results).To(ContainElement(failure))
			Expect(actualResult.Results).To(ContainElement(error))
		})

		It("should parse standard output from failure log", func() {
			// given
			filepath := getFilepath("multiple_test_failures_junit_surefire.xml")
			expectedSummary := ExecutionSummary{
				Total: 22, Failures: 2, Errors: 1, Skipped: 7, Time: "0.055", SystemOut: "CubeDockerConfiguration:\n  serverUri = tcp://localhost:4243\n  tlsVerify = false\n  dockerServerIp = localhost\n  definitionFormat = COMPOSE\n  clean = false\n  removeVolumes = true\n  dockerContainers = containers: {}\nnetworks: {}\n\n\nCubeDockerConfiguration:\n  serverUri = tcp://localhost:4243\n  tlsVerify = false\n  dockerServerIp = localhost\n  definitionFormat = COMPOSE\n  clean = false\n  removeVolumes = true\n  dockerContainers = containers: {}\nnetworks: {}\n\n\n", SystemErr: "",
			}
			surefireParser := SurefireParser{}

			// when
			actualResult, _ := surefireParser.Parse(filepath)

			// then
			Expect(actualResult.Summary).To(Equal(expectedSummary))
		})

		It("should convert entire report", func() {
			// given
			filepath := getFilepath("multiple_test_failures_testng_surefire.xml")
			results := []TestResult{{
				TestCase: "bar",
				Time:     "0.021",
				Kind:     FAILURE,
				Type:     "java.lang.AssertionError",
				Message:  "expected [foo] but found [2.0]",
				Details:  "java.lang.AssertionError: expected [foo] but found [2.0]\n\tat org.testng.Assert.fail(Assert.java:94)\n\tat org.testng.Assert.failNotEquals(Assert.java:494)\n\tat org.testng.Assert.assertEquals(Assert.java:123)\n\tat org.testng.Assert.assertEquals(Assert.java:165)\n\tat testNgMavenExample1.TestNgMavenExampleTest.bar(TestNgMavenExampleTest.java:21)\n",
			},
				{
					TestCase: "exampleOfTestNgMaven",
					Time:     "0",
					Kind:     PASSED,
					Message:  "",
					Details:  "",
				},
				{
					TestCase: "foo",
					Time:     "0.005",
					Kind:     FAILURE,
					Type:     "java.lang.AssertionError",
					Message:  "expected [foo] but found [1.0]",
					Details:  "java.lang.AssertionError: expected [foo] but found [1.0]\n\tat org.testng.Assert.fail(Assert.java:94)\n\tat org.testng.Assert.failNotEquals(Assert.java:494)\n\tat org.testng.Assert.assertEquals(Assert.java:123)\n\tat org.testng.Assert.assertEquals(Assert.java:165)\n\tat testNgMavenExample1.TestNgMavenExampleTest.foo(TestNgMavenExampleTest.java:16)\n",
				},
				{
					TestCase: "testCaseSkipException",
					Time:     "0.002",
					Kind:     SKIPPED,
					Message:  "",
					Details:  "",
				},
			}
			expectedResult := &TestResults{
				Name: "testNgMavenExample1.TestNgMavenExampleTest",
				Summary: ExecutionSummary{
					Total: 4, Failures: 2, Errors: 0, Skipped: 1, Time: "0.028", SystemOut: "Configuring TestNG with: org.apache.maven.surefire.testng.conf.TestNG652Configurator@3d82c5f3\n    Im in skip exception\n    ",
					SystemErr: "",
				},
				Results: results}

			surefireParser := SurefireParser{}

			// when
			actualResult, _ := surefireParser.Parse(filepath)

			// then
			Expect(actualResult).To(Equal(expectedResult))
		})
	})

	Describe("Parsing Maven Failsafe summary report", func() {
		It("Should convert summary report to TestResult", func() {
			// given
			filepath := getFilepath("failsafe-summary.xml")
			expectedResult := &TestResults{
				Name: "254",
				Summary: ExecutionSummary{
					Total: 0, Failures: 1, Errors: 0, Skipped: 0, SystemOut: "", SystemErr: "",
				},
				Results: nil}
			failsafeSummaryParser := FailsafeParser{}

			// when
			actualResult, _ := failsafeSummaryParser.Parse(filepath)

			// then
			Expect(actualResult).To(Equal(expectedResult))
		})
	})

	// Gradle and Surefire reports are based on the same JUnit XML schema
	Describe("Parsing Gradle default test results", func() {
		It("Should convert single test error to TestResult", func() {
			// given
			filepath := getFilepath("gradle_single_test_error.xml")
			f := []TestResult{{
				TestCase: "test",
				Time:     "0.003",
				Kind:     FAILURE,
				Message:  "java.lang.AssertionError",
				Type:     "java.lang.AssertionError",
				Details:  "java.lang.AssertionError\n\tat com.udacity.gradle.test.PersonTest.test(PersonTest.java:10)\n\tat sun.reflect.NativeMethodAccessorImpl.invoke0(Native Method)\n\tat sun.reflect.NativeMethodAccessorImpl.invoke(NativeMethodAccessorImpl.java:62)\n\tat sun.reflect.DelegatingMethodAccessorImpl.invoke(DelegatingMethodAccessorImpl.java:43)\n\tat java.lang.reflect.Method.invoke(Method.java:497)\n\tat org.junit.runners.model.FrameworkMethod$1.runReflectiveCall(FrameworkMethod.java:50)\n\tat org.junit.internal.runners.model.ReflectiveCallable.run(ReflectiveCallable.java:12)\n\tat org.junit.runners.model.FrameworkMethod.invokeExplosively(FrameworkMethod.java:47)\n\tat org.junit.internal.runners.statements.InvokeMethod.evaluate(InvokeMethod.java:17)\n\tat org.junit.runners.ParentRunner.runLeaf(ParentRunner.java:325)\n\tat org.junit.runners.BlockJUnit4ClassRunner.runChild(BlockJUnit4ClassRunner.java:78)\n\tat org.junit.runners.BlockJUnit4ClassRunner.runChild(BlockJUnit4ClassRunner.java:57)\n\tat org.junit.runners.ParentRunner$3.run(ParentRunner.java:290)\n\tat org.junit.runners.ParentRunner$1.schedule(ParentRunner.java:71)\n\tat org.junit.runners.ParentRunner.runChildren(ParentRunner.java:288)\n\tat org.junit.runners.ParentRunner.access$000(ParentRunner.java:58)\n\tat org.junit.runners.ParentRunner$2.evaluate(ParentRunner.java:268)\n\tat org.junit.runners.ParentRunner.run(ParentRunner.java:363)\n\tat org.gradle.api.internal.tasks.testing.junit.JUnitTestClassExecuter.runTestClass(JUnitTestClassExecuter.java:112)\n\tat org.gradle.api.internal.tasks.testing.junit.JUnitTestClassExecuter.execute(JUnitTestClassExecuter.java:56)\n\tat org.gradle.api.internal.tasks.testing.junit.JUnitTestClassProcessor.processTestClass(JUnitTestClassProcessor.java:66)\n\tat org.gradle.api.internal.tasks.testing.SuiteTestClassProcessor.processTestClass(SuiteTestClassProcessor.java:51)\n\tat sun.reflect.NativeMethodAccessorImpl.invoke0(Native Method)\n\tat sun.reflect.NativeMethodAccessorImpl.invoke(NativeMethodAccessorImpl.java:62)\n\tat sun.reflect.DelegatingMethodAccessorImpl.invoke(DelegatingMethodAccessorImpl.java:43)\n\tat java.lang.reflect.Method.invoke(Method.java:497)\n\tat org.gradle.messaging.dispatch.ReflectionDispatch.dispatch(ReflectionDispatch.java:35)\n\tat org.gradle.messaging.dispatch.ReflectionDispatch.dispatch(ReflectionDispatch.java:24)\n\tat org.gradle.messaging.dispatch.ContextClassLoaderDispatch.dispatch(ContextClassLoaderDispatch.java:32)\n\tat org.gradle.messaging.dispatch.ProxyDispatchAdapter$DispatchingInvocationHandler.invoke(ProxyDispatchAdapter.java:93)\n\tat com.sun.proxy.$Proxy2.processTestClass(Unknown Source)\n\tat org.gradle.api.internal.tasks.testing.worker.TestWorker.processTestClass(TestWorker.java:109)\n\tat sun.reflect.NativeMethodAccessorImpl.invoke0(Native Method)\n\tat sun.reflect.NativeMethodAccessorImpl.invoke(NativeMethodAccessorImpl.java:62)\n\tat sun.reflect.DelegatingMethodAccessorImpl.invoke(DelegatingMethodAccessorImpl.java:43)\n\tat java.lang.reflect.Method.invoke(Method.java:497)\n\tat org.gradle.messaging.dispatch.ReflectionDispatch.dispatch(ReflectionDispatch.java:35)\n\tat org.gradle.messaging.dispatch.ReflectionDispatch.dispatch(ReflectionDispatch.java:24)\n\tat org.gradle.messaging.remote.internal.hub.MessageHub$Handler.run(MessageHub.java:360)\n\tat org.gradle.internal.concurrent.ExecutorPolicy$CatchAndRecordFailures.onExecute(ExecutorPolicy.java:54)\n\tat org.gradle.internal.concurrent.StoppableExecutorImpl$1.run(StoppableExecutorImpl.java:40)\n\tat java.util.concurrent.ThreadPoolExecutor.runWorker(ThreadPoolExecutor.java:1142)\n\tat java.util.concurrent.ThreadPoolExecutor$Worker.run(ThreadPoolExecutor.java:617)\n\tat java.lang.Thread.run(Thread.java:745)\n"}}

			expectedResult := &TestResults{
				Name: "com.udacity.gradle.test.PersonTest",
				Summary: ExecutionSummary{
					Total: 1, Failures: 1, Errors: 0, Skipped: 0, Time: "0.003", SystemOut: "", SystemErr: "",
				},
				Results: f}
			gradleParser := GradleParser{}

			// when
			actualResult, _ := gradleParser.Parse(filepath)

			//then
			Expect(actualResult).To(Equal(expectedResult))
		})
	})

})

// --- Test helpers

func getFilepath(filename string) string {
	_, curr_dir, _, ok := runtime.Caller(0)
	if !ok {
		Fail("Unable to find path using runtime.Caller")
	}
	targetDir := filepath.FromSlash(path.Dir(curr_dir) + "/test_fixtures/")
	return filepath.FromSlash(targetDir + filename)
}
