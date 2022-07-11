package did_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestDid(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Did Suite")
}
