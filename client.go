package client

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

// MyFxBookClient struct
type MyFxBookClient struct {
	email    string
	password string
	session  string

	httpClient *http.Client
	loc        *time.Location
}

// NewClient creates new myfxbook client
func NewClient(email, password string, loc *time.Location) *MyFxBookClient {
	return &MyFxBookClient{
		email:      email,
		password:   password,
		httpClient: &http.Client{},
		loc:        loc,
	}
}

// NewWithHTTPClient creates new myfxbook client by a custom HTTP client
func NewWithHTTPClient(email, password string, loc *time.Location, client *http.Client) *MyFxBookClient {
	return &MyFxBookClient{
		email:      email,
		password:   password,
		httpClient: client,
		loc:        loc,
	}
}

func (c *MyFxBookClient) request(method, url string) ([]byte, error) {
	var (
		err      error
		response []byte
	)

	switch method {
	case http.MethodGet:
		response, err = c.get(url)
	case http.MethodPost:
		response, err = c.post(url, nil)
	}

	return response, err
}

func (c *MyFxBookClient) get(url string) ([]byte, error) {
	response, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}

func (c *MyFxBookClient) post(url string, payload interface{}) ([]byte, error) {
	response, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}

// Login to account
func (c *MyFxBookClient) Login() (*LoginResponse, error) {
	var err error

	data, err := c.request(http.MethodGet, getLoginURL(c.email, c.password))
	if err != nil {
		return nil, err
	}

	response := &LoginResponse{}
	err = json.Unmarshal(data, response)
	c.session = response.Session
	return response, err
}

// Logout from account
func (c *MyFxBookClient) Logout() (*Response, error) {
	var err error

	data, err := c.request(http.MethodGet, getLogoutURL(c.session))
	if err != nil {
		return nil, err
	}

	response := &Response{}
	err = json.Unmarshal(data, response)
	return response, err
}

// GetMyAccounts get a list of my accounts and their data
func (c *MyFxBookClient) GetMyAccounts() (*GetMyAccountsResponse, error) {
	var err error

	data, err := c.request(http.MethodGet, getGetMyAccountsURL(c.session))
	if err != nil {
		return nil, err
	}

	response := &GetMyAccountsResponse{}
	err = json.Unmarshal(data, response)
	return response, err
}
