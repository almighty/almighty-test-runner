package buildtool_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDetectors(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Build Tool Detectors Suite")
}
