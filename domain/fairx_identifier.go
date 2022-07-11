package domain

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
)

const (
	DID_SEPARATOR    = ":"
	DID_METHOD_FAIRX = "fairx"
)

// target@domain.com:/path
type FairXIdentifier struct {
	Target string
	Domain string
	Path   string
}

func (f *FairXIdentifier) DID() string {

	rawdidid := fmt.Sprintf("%s@%s", f.Target, f.Domain)
	if f.Path != "" {
		rawdidid = fmt.Sprintf("%s:%s", rawdidid, f.Path)
	}

	encoded := base64.RawURLEncoding.EncodeToString([]byte(rawdidid))
	return fmt.Sprintf("did:fairx:%s", encoded)

}

func ParseDIDIdentifier(did string) (*FairXIdentifier, error) {

	// split into the three required parts
	parts := strings.Split(did, DID_SEPARATOR)
	if len(parts) != 3 || parts[0] != "did" {
		return nil, errors.New("Invalid fairx DID")
	}

	// Make sure the 2nd component is "fairx"
	if parts[1] != DID_METHOD_FAIRX {
		return nil, errors.New("DID method is not fairx")
	}

	// Base64url decode the 3rd part
	if len(parts[2]) == 0 {
		return nil, errors.New("FairX DID method requires an identifier")
	}

	identBytes, err := base64.RawURLEncoding.DecodeString(parts[2])
	if err != nil {
		return nil, errors.New("FairX DID method identifier could not be base64 url decoded")
	}

	// parse the ident
	fairxIdentifier := FairXIdentifier{}
	identParts := strings.Split(string(identBytes), "@")
	if len(identParts) != 2 {
		return nil, errors.New("FairX Identifier was not properly formatted.  Must be ident@domain.com:/optional/path")
	}
	fairxIdentifier.Target = identParts[0]

	// parse the optional path, if any
	if strings.Index(identParts[1], ":") > 0 {
		pathParts := strings.Split(identParts[1], ":")
		fairxIdentifier.Domain = pathParts[0]
		fairxIdentifier.Path = pathParts[1]
	} else {
		fairxIdentifier.Domain = identParts[1]
	}

	return &fairxIdentifier, nil

}
