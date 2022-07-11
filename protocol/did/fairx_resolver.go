package did

import (
	"github.com/fairxio/go/did"
)

type FairXDIDResolver struct {
}

func CreateFairXDIDResolver() *FairXDIDResolver {
	return &FairXDIDResolver{}
}

func (r *FairXDIDResolver) Resolve(did string) did.DIDDocument {
	//TODO implement me
	panic("implement me")
}
