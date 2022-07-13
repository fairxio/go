package didsrv

import (
	"github.com/fairxio/go/applications/fairx/configuration"
	"github.com/fairxio/go/log"
)

func Start() {

	// Create Configuration
	config := configuration.Create()

	httpService := CreateHTTPService(config.GetListenAddress(), config.GetListenPort())
	log.Info("FairX DID Service shutting down:  %v", httpService.ListenAndServe())

}
