package ovn

import (
	"testing"

	 "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestClusterNode(t *testing.T) {
	RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "OVN Operations Suite")
}
