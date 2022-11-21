package sdk

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

// Options - required for a new instance of SDK Core Client
type Options struct {
	BaseURI     string
	AccessToken string
	APIVersion  string
	// QueryParams map[string]string
	Debug bool
}

// PrettyPrintJSON takes an raw interface and prints json to stdout
func PrettyPrintJSON(data interface{}) {
	resultBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatal("failed to pretty print json string:", err)
	}

	fmt.Println(string(resultBytes))
}

// InitClient - initializes a new instance of ADO Core client
func InitClient(options *Options) *resty.Request {
	client := resty.New()
	return client.R().
		EnableTrace().
		// SetQueryParams(options.QueryParams).
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", options.AccessToken)
}

// Get requests raw
func (t *Options) Get(path string, options *Options) (*resty.Response, error) {
	client := InitClient(t)
	return client.Get(t.BaseURI + path)
}

// GetI get as interface
func (t *Options) GetI(path string, res interface{}) (*resty.Response, error) {
	client := InitClient(t)
	return client.SetResult(res).Get(t.BaseURI + path)
}

// Post requests
func (t *Options) Post(path string, body interface{}) (*resty.Response, error) {
	client := InitClient(t)
	return client.SetBody(body).Post(t.BaseURI + path)
}

// PostI requests
func (t *Options) PostI(path string, body interface{}, res interface{}) (*resty.Response, error) {
	client := InitClient(t)
	return client.SetResult(res).SetBody(body).Post(t.BaseURI + path)
}

// Put requests
func (t *Options) Put(path string, body interface{}) (*resty.Response, error) {
	client := InitClient(t)
	return client.SetBody(body).Put(t.BaseURI + path)
}

// PutI requests
func (t *Options) PutI(path string, body interface{}, res interface{}) (*resty.Response, error) {
	client := InitClient(t)
	return client.SetResult(res).SetBody(body).Put(t.BaseURI + path)
}

// Delete requests
func (t *Options) Delete(path string, body interface{}) (*resty.Response, error) {
	client := InitClient(t)
	return client.SetBody(body).Delete(t.BaseURI + path)
}

// DeleteI requests
func (t *Options) DeleteI(path string, body interface{}, res interface{}) (*resty.Response, error) {
	client := InitClient(t)
	return client.SetResult(res).SetBody(body).Put(t.BaseURI + path)
}
