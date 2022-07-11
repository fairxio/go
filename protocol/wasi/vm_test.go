package wasi_test

import (
	"github.com/fairxio/go/protocol/wasi"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"io/ioutil"
)

var _ = Describe("VM", func() {

	Describe("Loading a WASI Package", func() {

		var virtualMachine *wasi.VirtualMachine
		var jsBytes []byte
		var err error

		BeforeEach(func() {
			jsBytes, err = ioutil.ReadFile("applications/wasm/example/js_simple_example.js")
			Expect(err).To(BeNil())
		})

		It("Loads the WASM and validates", func() {

			virtualMachine = wasi.CreateVirtualMachine(jsBytes)
			Expect(virtualMachine).ToNot(BeNil())

		})

		It("Executes a simple function", func() {

			err = virtualMachine.ExecuteFunction("PerformExecutableWorkflow")
			Expect(err).To(BeNil())

		})

		It("Provides a function in the VM to the Javascript", func() {

			virtualMachine.ProvideFunction("callParticipant", virtualMachine.CallParticipant)
			virtualMachine.ExecuteFunction("PerformParticipantCalling", "participantIdentity")

		})

	})

})
