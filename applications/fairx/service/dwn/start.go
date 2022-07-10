package dwn

import (
	"github.com/fairxio/go/applications/fairx/configuration"
	"github.com/fairxio/go/log"
)

func Start() {

	// Get Configuration
	config := configuration.Create()

	// Create HTTP Service
	httpService := CreateHTTPService(config.GetListenAddress(), config.GetListenPort())

	// Wait and Serve
	log.Fatal("Stopped Listening and Serving:  %v", httpService.ListenAndServe())
}
