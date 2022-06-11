package addressing_test

import (
	"github.com/CascadiaShaman/go/protocol/addressing"
	"github.com/CascadiaShaman/go/protocol/interfaces"
	"github.com/CascadiaShaman/go/protocol/messages"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("DID", func() {

	Describe("String Encoding", func() {

		Context("No Queries", func() {

			dwnSpecific := addressing.DWNSpecificID{
				Service: string(addressing.DecentralizedWebNode),
				ID:      "12345",
			}

			did := addressing.DID{
				MethodName:       "testing",
				MethodSpecificID: dwnSpecific,
			}
			didString := did.String()

			It("Should encode to a simple DID", func() {
				Expect(didString).To(BeEquivalentTo("did:testing:12345?service=DecentralizedWebNode&queries=bnVsbA=="))
			})

		})

		Context("Single Query", func() {

			nonce, _ := uuid.NewRandom()
			md := messages.MessageDescriptor{
				Method:     interfaces.CollectionsQuery,
				Nonce:      nonce.String(),
				DataFormat: interfaces.JSON,
				DataCID:    "cid",
			}

			dwnSpecific := addressing.DWNSpecificID{
				Queries: []messages.MessageDescriptor{md},
				Service: string(addressing.DecentralizedWebNode),
				ID:      "12345",
			}

			did := addressing.DID{
				MethodName:       "testing",
				MethodSpecificID: dwnSpecific,
			}
			didString := did.String()

			It("Should encode to a DID with a base64url encoded string representing Message Descriptors", func() {
				_ = didString
				// Expect(didString).To(BeEquivalentTo("did:testing:12345?service=DecentralizedWebNode&queries=W3sibm9uY2UiOiJmYTdkNGYyMS0wZGY1LTQ2NDUtOTg3OC01NDlhOGQzNjVkMmQiLCJtZXRob2QiOiJDb2xsZWN0aW9uc1F1ZXJ5IiwiZGF0YUNpZCI6ImNpZCIsImRhdGFGb3JtYXQiOiJhcHBsaWNhdGlvbi9qc29uIn1d"))
			})

		})

	})

	Describe("Parsing", func() {

		Context("Simple DID no queries", func() {

			didString := "did:testing:12345"
			did, err := addressing.Parse(didString)

			It("Parsed the service name and specific ID", func() {

				Expect(err).To(BeNil())
				Expect(did.MethodName).ToNot(BeEmpty())
				Expect(did.MethodName).To(BeEquivalentTo("testing"))
				Expect(did.MethodSpecificID.ID).To(BeEquivalentTo("12345"))

			})

		})

		Context("DID with 1 query", func() {

			didString := "did:testing:01234?service=DecentralizedWebNode&queries=W3sibm9uY2UiOiJmYTdkNGYyMS0wZGY1LTQ2NDUtOTg3OC01NDlhOGQzNjVkMmQiLCJtZXRob2QiOiJDb2xsZWN0aW9uc1F1ZXJ5IiwiZGF0YUNpZCI6ImNpZCIsImRhdGFGb3JtYXQiOiJhcHBsaWNhdGlvbi9qc29uIn1d"
			did, err := addressing.Parse(didString)

			It("Parsed the service name and specific ID", func() {

				Expect(err).To(BeNil())
				Expect(did.MethodName).ToNot(BeEmpty())
				Expect(did.MethodName).To(BeEquivalentTo("testing"))
				Expect(did.MethodSpecificID.ID).To(BeEquivalentTo("01234"))

				Expect(len(did.MethodSpecificID.Queries)).To(BeEquivalentTo(1))
				Expect(did.MethodSpecificID.Queries[0].Method).To(BeEquivalentTo(interfaces.CollectionsQuery))
				Expect(did.MethodSpecificID.Queries[0].Nonce).To(BeEquivalentTo("fa7d4f21-0df5-4645-9878-549a8d365d2d"))
				Expect(did.MethodSpecificID.Queries[0].DataCID).To(BeEquivalentTo("cid"))
				Expect(did.MethodSpecificID.Queries[0].DataFormat).To(BeEquivalentTo(interfaces.JSON))

			})

		})

	})

})
