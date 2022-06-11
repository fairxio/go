package addressing

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/CascadiaShaman/go/protocol/messages"
	"net/url"
	"strings"
)

type Service string

const (
	DecentralizedWebNode Service = "DecentralizedWebNode"
)

// Specification: https://identity.foundation/decentralized-web-node/spec/#addressing
// The following process defines how a DID-Relative URL is composed to address a Decentralized Web Node:
//
// Let the base URI authority portion of the DID URL string be the target DID being addressed.
// Append a service parameter to the DID URL string with the value DecentralizedWebNode.
// Assemble an array of the Message Descriptor objects are desired for encoding in the DID-relative URL
// JSON stringify the array of Message Descriptor objects from Step 3, then Base64Url encode the stringified output.
// Append a queries parameter to the DID URL string with the value set to the JSON stringified, Base64Url encoded output of Step 4.
// DID-relative URLs are composed of the following segments
//
// did:example:123 + ?service=DecentralizedWebNode + &queries= + toBase64Url( JSON.stringify( [{ DESCRIPTOR_1 }, { DESCRIPTOR_N }] ) )
type DID struct {
	MethodName       string
	MethodSpecificID DWNSpecificID
}

func (did DID) String() string {
	return fmt.Sprintf("did:%s:%s", did.MethodName, did.MethodSpecificID.String())
}

type DWNSpecificID struct {
	ID      string
	Service string
	Queries []messages.MessageDescriptor
}

func (dwn DWNSpecificID) String() string {

	// Assemble Queries
	queryStringBytes, _ := json.Marshal(dwn.Queries)
	queryString := base64.URLEncoding.EncodeToString(queryStringBytes)

	specificId := fmt.Sprintf("%s?service=%s&queries=%s", dwn.ID, dwn.Service, queryString)
	return specificId
}

func Parse(didString string) (DID, error) {

	didPieces := strings.Split(didString, ":")
	if len(didPieces) != 3 {
		return DID{}, errors.New("Unrecognized DID URL Format")
	}

	questionIndex := strings.Index(didPieces[2], "?")
	if questionIndex == -1 {

		methodSpec := DWNSpecificID{ID: didPieces[2]}
		return DID{
			MethodName:       didPieces[1],
			MethodSpecificID: methodSpec,
		}, nil
	}

	dwn := DWNSpecificID{}
	if questionIndex > 0 {

		// get what is in between
		dwn.ID = didPieces[2][0:questionIndex]
	}

	queryString := didPieces[2][questionIndex+1:]
	q, err := url.ParseQuery(queryString)
	if err != nil {
		return DID{}, err
	}
	dwn.Service = q.Get("service")

	// Base64URL Decode
	queriesJsonString, err := base64.URLEncoding.DecodeString(q.Get("queries"))
	if err != nil {
		return DID{}, err
	}

	var queries []messages.MessageDescriptor
	err = json.Unmarshal(queriesJsonString, &queries)
	if err != nil {
		return DID{}, err
	}

	if len(queries) > 0 {
		dwn.Queries = queries
	}

	did := DID{
		MethodName:       didPieces[1],
		MethodSpecificID: dwn,
	}

	return did, nil

}
