package ninja

import (
	"os"
	"fmt"
	"net/http"
)

type Client struct {
	client	*http.Client
	token	string
	baseURL	string
}
func NewClient() *Client {
	client := &http.Client{}

	return &Client{
		client: client,
		token: GetCred(client), 
		baseURL: os.Getenv("NINJAONE_URL"),
	}
}

func (c *Client) SetHeader(req *http.Request) {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "Kawika's CLI Stuff")
}
