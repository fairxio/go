package auth

import (
	"encoding/json"
	"fmt"
	"github.com/fairxio/go/applications/fairx/configuration"
	jwt "github.com/fairxio/go/authentication/jwt"
	"github.com/fairxio/go/ext"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

type HTTPService struct {
	ListenAddr string
	ListenPort int
	Router     *mux.Router

	NonceCache ext.CacheService
}

func CreateHTTPService(addr string, port int) *HTTPService {

	svc := HTTPService{}
	svc.Router = mux.NewRouter()
	svc.ListenAddr = addr
	svc.ListenPort = port
	svc.Router.HandleFunc("/v1.0.0/auth", svc.ServiceHandler).Methods(http.MethodPost)

	svc.NonceCache = CreateSimpleNonceCache()

	return &svc

}

func (svc *HTTPService) ListenAndServe() error {
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", svc.ListenAddr, svc.ListenPort), svc.Router)
	return err
}

func (svc *HTTPService) ServiceHandler(w http.ResponseWriter, r *http.Request) {

	// Check Authentication Request
	body, err := ioutil.ReadAll(r.Body)
	if err != nil || body == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var authReq AuthenticationRequest
	err = json.Unmarshal(body, &authReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if authReq.ID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// If we have just an ID and no sig, return with a nonce
	config := configuration.Create()
	if authReq.Signature == "" {

		// Generate and cache a nonce for this DID
		nonceUUID := uuid.New().String()
		svc.NonceCache.Put(authReq.ID, []byte(nonceUUID))

		// Return with nonce
		authReq.Nonce = nonceUUID
		authReqBytes, _ := json.Marshal(&authReq)

		// reply to client
		w.Header().Set("Content-Type", "application/json")
		w.Write(authReqBytes)
		return

	} else {

		// DOnt even bother verifying the signature for now
		// TODO:  Verify signature
		token := jwt.GenerateJWT(authReq.ID, "did:fairx:issuerDID")
		signedToken := jwt.SignJWT(token, config.GetJWTKey())

		authReq.JWT = signedToken
		resp, _ := json.Marshal(&authReq)
		w.Header().Set("Content-Type", "application/json")
		w.Write(resp)
		return

	}

}
