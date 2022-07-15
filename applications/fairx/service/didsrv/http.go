package didsrv

import (
	"fmt"
	"github.com/fairxio/go/authentication/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

type HTTPService struct {
	ListenAddr string
	ListenPort int
	Router     *mux.Router
}

func CreateHTTPService(addr string, port int, jwtKey string) *HTTPService {

	middleware.JWTSecretKey = jwtKey
	svc := HTTPService{}
	svc.Router = mux.NewRouter()
	svc.ListenAddr = addr
	svc.ListenPort = port
	svc.Router.HandleFunc("/v1.0.0", middleware.IsJWTAuthorized(svc.ServiceHandler)).Methods(http.MethodPost)

	return &svc

}

func (svc *HTTPService) ListenAndServe() error {
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", svc.ListenAddr, svc.ListenPort), svc.Router)
	return err
}

func (svc *HTTPService) ServiceHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
