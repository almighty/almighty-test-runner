package testresultparser_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTestResultsParser(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test Results parsing Suite")
}
