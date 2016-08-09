package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestAlmightyTestRunner(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AlmightyTestRunner Suite")
}
