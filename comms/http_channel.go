package comms

import (
	"io/ioutil"
	"net/http"
)

type GoHTTPChannel struct{}

func CreateGoHTTPChannel() GoHTTPChannel {
	return GoHTTPChannel{}
}

func (channel GoHTTPChannel) Get(url string) ([]byte, error) {

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil

}
