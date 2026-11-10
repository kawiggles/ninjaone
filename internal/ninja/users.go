package ninja

import (
	"encoding/json"
	"net/http"
	"io"
)

func UnmarshalUsers(data []byte) ([]User, error) {
	var r []User
	err := json.Unmarshal(data, &r)
	return r, err
}

type User struct {
	ID                int64    `json:"id"`
	Uid               string   `json:"uid"`
	FirstName         string   `json:"firstName"`
	LastName          string   `json:"lastName"`
	Email             string   `json:"email"`
	Phone             string   `json:"phone"`
	Enabled           bool     `json:"enabled"`
	Administrator     bool     `json:"administrator"`
	PermitAllClients  bool     `json:"permitAllClients"`
	NotifyAllClients  bool     `json:"notifyAllClients"`
	MustChangePw      bool     `json:"mustChangePw"`
	MfaConfigured     bool     `json:"mfaConfigured"`
	UserType          string   `json:"userType"`
	InvitationStatus  string   `json:"invitationStatus"`
	OrganizationID    int64    `json:"organizationId"`
	DeviceIDS         []int64  `json:"deviceIds"`
	AssignedDeviceIDS []int64  `json:"assignedDeviceIds"`
	Roles             []string `json:"roles"`
}

func (c *Client) GetUsers() ([]User, error) {
	resp, err := http.Get(c.baseURL+"/v2/users")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	users, err := UnmarshalUsers(body)
	if err != nil {
		return nil, err
	}

	return users, nil
}
