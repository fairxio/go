package domain_test

import (
	"github.com/fairxio/go/domain"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Parts", func() {

	testDid := "did:fairx:dGVzdGluZ0BmYWlyeC5pbzovc29tZS90ZXN0L3BhdGgjd2l0aHRhZ3M"
	testDidNoPath := "did:fairx:dGVzdGluZ0BmYWlyeC5pbw"
	testFairXIdentifier := domain.FairXIdentifier{
		Target: "testing",
		Domain: "fairx.io",
		Path:   "/some/test/path#withtags",
	}
	testFairXIdentifierNoPath := domain.FairXIdentifier{
		Target: testFairXIdentifier.Target,
		Domain: testFairXIdentifier.Domain,
	}

	Describe("Generating a Valid DID String", func() {

		didId := testFairXIdentifier.DID()

		It("Created a valid FairX did: identifier", func() {
			Expect(didId).ToNot(BeNil())
			Expect(didId).To(Equal(testDid))
		})

	})

	Describe("Generating a Valid DID String without a path component", func() {

		didId := testFairXIdentifierNoPath.DID()

		It("Created a valid FairX did: identifier", func() {
			Expect(didId).ToNot(BeNil())
			Expect(didId).To(Equal(testDidNoPath))
		})

	})

	Describe("Parsing a DID String", func() {

		ident, err := domain.ParseDIDIdentifier(testDid)

		It("Was able to parse a valid DID string", func() {
			Expect(err).To(BeNil())
			Expect(ident.Target).To(Equal(testFairXIdentifier.Target))
			Expect(ident.Domain).To(Equal(testFairXIdentifier.Domain))
			Expect(ident.Path).To(Equal(testFairXIdentifier.Path))
		})

	})

	Describe("Parsing a DID string without a path component", func() {

		ident, err := domain.ParseDIDIdentifier(testDidNoPath)

		It("Was able to parse a valid DID string without a path component", func() {
			Expect(err).To(BeNil())
			Expect(ident.Target).To(Equal(testFairXIdentifier.Target))
			Expect(ident.Domain).To(Equal(testFairXIdentifier.Domain))
			Expect(ident.Path).To(BeEmpty())
		})

	})

})
