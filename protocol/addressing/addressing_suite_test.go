package addressing_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestAddressing(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Addressing Suite")
}
