package auth

type AuthenticationRequest struct {
	ID        string `json:"id"`
	Nonce     string `json:"nonce,omitempty"`
	Signature string `json:"sig,omitempty"`
	JWT       string `json:"jwt,omitempty"`
}
