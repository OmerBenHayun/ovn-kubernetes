package util

import (
	"testing"

	 "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestUtilSuite(t *testing.T) {
	RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Util Suite")
}
