package did

type DIDDocument struct {
	ID                   string               `json:"id"`
	AlsoKnownAs          []string             `json:"alsoKnownAs"`
	Controller           []string             `json:"controller"`
	VerificationMethods  []VerificationMethod `json:"verificationMethod"`
	Authentication       []VerificationMethod `json:"authentication"`
	AssertionMethod      []VerificationMethod `json:"assertionMethod"`
	KeyAgreement         []VerificationMethod `json:"keyAgreement"`
	CapabilityInvocation []VerificationMethod `json:"capabilityInvocation"`
	CapabilityDelegation []VerificationMethod `json:"capabilityDelegation"`
	Service              []Service            `json:"service"`
}

type VerificationMethod struct {
	ID         string            `json:"id"`
	Type       string            `json:"type"`
	Controller string            `json:"controller"`
	Methods    map[string]string `json:"methods"`
}

type ServiceEndpoint struct {
	ID      string `json:"id"`
	Purpose string `json:"purpose"`
	URI     string `json:"uri"`
}

type Service struct {
	ID        string            `json:"id"`
	Endpoints []ServiceEndpoint `json:"endpoints"`
}
