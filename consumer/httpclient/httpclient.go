package httpclient

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	httpClientTimeout = time.Minute
)

//Client to handle http requests
type Client struct {
	*http.Client

	username string
	password string
	log      *logrus.Entry
}

//NewHTTPClient returns new Client object
func NewHTTPClient(username string, password string,
	httpTransport *http.Transport, log *logrus.Entry) *Client {

	return &Client{
		Client: &http.Client{
			Transport: httpTransport,
			Timeout:   httpClientTimeout,
		},
		username: username,
		password: password,
		log:      log,
	}

}

//Do method
func (cl *Client) Do(ctx context.Context, method string,
	url string, requestBody []byte) ([]byte, int, error) {

	log := cl.log

	req, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		log.Error(err)
		return nil, 0, err
	}

	// if cl.username != "" && cl.password != "" {
	// 	log.Infof("Using BasicAuth for authentication with the Server: %s", url)
	// 	req.SetBasicAuth(cl.username, cl.password)
	// }

	req.SetBasicAuth(cl.username, cl.password)

	req.Header = http.Header{}
	req.Header["Content-Type"] = []string{"application/json"}

	log.Infof(">>" + method + " " + url + " " + string(requestBody))

	resp, err := cl.Client.Do(req)
	if err != nil {
		log.Error(err)
		return nil, 0, err
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		return nil, resp.StatusCode, err
	}

	log.Infof("status: %v response: %v", resp.StatusCode, string(respBody))

	return respBody, resp.StatusCode, nil

}
