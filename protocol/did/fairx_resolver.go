package did

import (
	"encoding/json"
	"fmt"
	"github.com/fairxio/go/comms"
	"github.com/fairxio/go/did"
	"github.com/fairxio/go/domain"
)

type FairXDIDResolver struct {
	Channel comms.Channel
}

func CreateFairXDIDResolver(channel comms.Channel) *FairXDIDResolver {
	return &FairXDIDResolver{
		Channel: channel,
	}
}

func (r *FairXDIDResolver) Resolve(didIdent string) *did.DIDDocument {

	// Parse DID
	fxId, err := domain.ParseDIDIdentifier(didIdent)
	if err != nil {
		return nil
	}

	// Create URL
	url := fmt.Sprintf("https://%s/v1.0.0/%s", fxId.Domain, didIdent)
	body, err := r.Channel.Get(url)

	// Parse the DID Document
	var didDoc did.DIDDocument
	err = json.Unmarshal(body, &didDoc)
	if err != nil {
		return nil
	}

	return &didDoc

}
