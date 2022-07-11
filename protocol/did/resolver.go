package did

import (
	"github.com/fairxio/go/comms"
	"github.com/fairxio/go/did"
)

type DIDResolver interface {
	Resolve(didIdent string) *did.DIDDocument
}

const (
	DID_METHOD_FAIRX = "fairx"
	DID_METHOD_KEY   = "key"
)

func ResolverForMethod(methodName string) DIDResolver {

	switch methodName {

	case DID_METHOD_FAIRX:
		httpChannel := comms.CreateGoHTTPChannel()
		return CreateFairXDIDResolver(httpChannel)

	case DID_METHOD_KEY:
		return CreateKeyDIDResolver()

	default:
		return nil

	}

}
