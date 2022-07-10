package did

import (
	"encoding/base64"
	"github.com/fairxio/go/log"
)

const (
	DATA_FORMAT_JSON   = "application/json"
	DATA_FORMAT_VC_JWT = "application/vc+jwt"
	DATA_FORMAT_VC_LDP = "application/vc+ldp"
)

type RequestObject struct {
	Target   string    `json:"target"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Data          string      `json:"data,omitempty"`
	Descriptor    Descriptor  `json:"descriptor,omitempty"`
	Attestation   interface{} `json:"attestation,omitempty"`
	Authorization interface{} `json:"authorization,omitempty"`
}

type Descriptor struct {
	Nonce      string `json:"nonce"`
	Method     string `json:"method"`
	DataCID    string `json:"dataCid"`
	DataFormat string `json:"dataFormat"`
}

type JWS struct {
	Payload    string         `json:"payload"`
	Signatures []JWSSignature `json:"signatures"`
}

type JWSSignature struct {
	Protected string `json:"protected"`
	Signature string `json:"signature"`
}

func (msg Message) GetDecodedData() []byte {

	var data []byte
	if msg.Data != "" {
		d, err := base64.URLEncoding.DecodeString(msg.Data)
		if err != nil {
			log.Warn("Unable to decode base64URL: %v", err)
			return nil
		}
		data = d
	}

	return data

}

func (msg Message) SetEncodedData(data []byte) {
	msg.Data = base64.URLEncoding.EncodeToString(data)
}
