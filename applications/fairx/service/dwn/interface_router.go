package dwn

import (
	"github.com/fairxio/go/did"
	"github.com/fairxio/go/log"
)

type InterfaceRouter struct {
	FeatureDetection
}

func CreateInterfaceRouter() *InterfaceRouter {

	ir := InterfaceRouter{
		FeatureDetection: CurrentFeatureDetection,
	}

	return &ir

}

func (ir *InterfaceRouter) Route(ro *did.RequestObject) (interface{}, error) {

	for _, msg := range ro.Messages {

		log.Info("Message: %v", msg)

	}

	return nil, nil

}

func (ir *InterfaceRouter) FeatureDetectionRoute() interface{} {

	return &CurrentFeatureDetection

}
