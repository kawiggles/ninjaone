package ninja

import (
	"encoding/json"
	"net/http"
)

type Device struct {
	id 			int32 `json:"id"`
	owneruid	string `json:"assignedOwnerUid"`
}

func (c *Client) GetDevices() ([]Device, error) {
	resp, err := http.Get(c.baseURL+"/v2/devices")
	if err != nil {
	}
}
