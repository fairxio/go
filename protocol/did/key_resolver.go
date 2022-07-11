package did

import (
	"github.com/fairxio/go/did"
)

type KeyDIDResolver struct {
}

func CreateKeyDIDResolver() *KeyDIDResolver {
	return &KeyDIDResolver{}
}

func (r *KeyDIDResolver) Resolve(did string) did.DIDDocument {
	//TODO implement me
	panic("implement me")
}
