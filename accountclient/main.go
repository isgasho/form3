package accountclient

import (
	
)

var (
	baseURL = "http://localhost:8080/v1/organisation/accounts"
)

// Client is the type used to interface with the Accounts API
type Client struct {
}

// New creates a new instance of an account client
func New() *Client {
	return &Client{}
}
