package did

import "github.com/fairxio/go/did"

type DIDResolver interface {
	Resolve(did string) did.DIDDocument
}

const (
	DID_METHOD_FAIRX = "fairx"
	DID_METHOD_KEY   = "key"
)

func ResolverForMethod(methodName string) DIDResolver {

	switch methodName {

	case DID_METHOD_FAIRX:
		return CreateFairXDIDResolver()

	case DID_METHOD_KEY:
		return CreateKeyDIDResolver()

	default:
		return nil

	}

}
