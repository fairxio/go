package dwn

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fairxio/go/authentication/middleware"
	"github.com/fairxio/go/did"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

type HTTPService struct {
	ListenAddr      string
	ListenPort      int
	Router          *mux.Router
	InterfaceRouter *InterfaceRouter
}

func CreateHTTPService(addr string, port int, jwtKey string) *HTTPService {

	// Set JWT Key
	middleware.JWTSecretKey = jwtKey
	svc := HTTPService{}
	svc.Router = mux.NewRouter()
	svc.InterfaceRouter = CreateInterfaceRouter()
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

	// Get Request Object
	ro, err := svc.GetRequestObject(w, r)
	if err != nil {
		featureDetectesponse := svc.InterfaceRouter.FeatureDetectionRoute()
		respBytes, err := json.Marshal(featureDetectesponse)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(respBytes)
		return

	}

	// Route Request
	routedResponse, err := svc.InterfaceRouter.Route(&ro)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// JSON Encode Response
	respBytes, err := json.Marshal(routedResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Respond to Client
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(respBytes)
	w.Header().Set("Content-Type", "application/json")

}

func (svc *HTTPService) GetRequestObject(w http.ResponseWriter, r *http.Request) (did.RequestObject, error) {

	var request did.RequestObject

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return request, errors.New("Bad Request")
	}

	err = json.Unmarshal(requestBody, &request)
	if err != nil {
		return request, errors.New("Non-JSON")
	}

	return request, nil

}
