package fxi_test

import (
	"github.com/fairxio/go/log"
	"github.com/fairxio/go/protocol/fxi"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"io/ioutil"
	v8 "rogchap.com/v8go"
)

var _ = Describe("VM", func() {

	Describe("Loading a WASI Package", func() {

		var virtualMachine *fxi.VirtualMachine
		var jsBytes []byte
		var err error

		BeforeEach(func() {
			jsBytes, err = ioutil.ReadFile("applications/wasm/example/js_simple_example.js")
			Expect(err).To(BeNil())
		})

		It("Loads the WASM and validates", func() {

			virtualMachine = fxi.CreateVirtualMachine(jsBytes)
			Expect(virtualMachine).ToNot(BeNil())

		})

		It("Executes a simple function", func() {

			err = virtualMachine.ExecuteFunction("PerformExecutableWorkflow")
			Expect(err).To(BeNil())

		})

		It("Provides a function in the VM to the Javascript", func() {

			virtualMachine.ProvideFunction("fairxTest", func(info *v8.FunctionCallbackInfo) *v8.Value {
				log.Info("FairX Test:  %v", info.Args())
				return nil
			})

			virtualMachine.ExecuteFunction("PerformParticipantCalling", "participantIdentity")

		})

	})

})
