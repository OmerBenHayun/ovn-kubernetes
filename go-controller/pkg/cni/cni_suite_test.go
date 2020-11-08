package cni

import (
	"testing"

	 "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCNISuite(t *testing.T) {
	RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "CNI Suite")
}
