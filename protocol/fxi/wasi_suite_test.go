package fxi_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestWasi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Wasi Suite")
}
