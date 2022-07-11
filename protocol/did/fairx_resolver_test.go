package did_test

import (
	"github.com/fairxio/go/mock"
	"github.com/fairxio/go/protocol/did"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("FairxResolver", func() {

	Describe("Resolving a valid FairX DID", func() {

		It("Should have returned a parsed DID Document", func() {
			ctrl := gomock.NewController(GinkgoT())

			mockChannel := mock.NewMockChannel(ctrl)
			mockChannel.EXPECT().Get("https://fairx.io/v1.0.0/did:fairx:dGVzdGluZ0BmYWlyeC5pbw").Return([]byte("{\"id\":\"did:fairx:dGVzdGluZ0BmYWlyeC5pbw\"}"), nil).Times(1)

			resolver := did.CreateFairXDIDResolver(mockChannel)
			didDoc := resolver.Resolve("did:fairx:dGVzdGluZ0BmYWlyeC5pbw")
			Expect(didDoc).ToNot(BeNil())
			Expect(didDoc.ID).To(Equal("did:fairx:dGVzdGluZ0BmYWlyeC5pbw"))

		})

	})

})
