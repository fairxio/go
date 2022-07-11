package domain

import (
	"encoding/base64"
	"fmt"
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
