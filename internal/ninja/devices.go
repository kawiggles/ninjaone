package ninja

import (
	"encoding/json"
	"io"
	"net/http"
)

func UnmarshalDevices(data []byte) ([]Device, error) {
	var r []Device
	err := json.Unmarshal(data, &r)
	return r, err
}

type Device struct {
	ID               int64       `json:"id"`
	Uid              string      `json:"uid"`
	AssignedOwnerUid string      `json:"assignedOwnerUid"`
	ParentDeviceID   int64       `json:"parentDeviceId"`
	OrganizationID   int64       `json:"organizationId"`
	LocationID       int64       `json:"locationId"`
	NodeClass        string      `json:"nodeClass"`
	NodeRoleID       int64       `json:"nodeRoleId"`
	RolePolicyID     int64       `json:"rolePolicyId"`
	PolicyID         int64       `json:"policyId"`
	ApprovalStatus   string      `json:"approvalStatus"`
	Offline          bool        `json:"offline"`
	DisplayName      string      `json:"displayName"`
	SystemName       string      `json:"systemName"`
	DNSName          string      `json:"dnsName"`
	NetbiosName      string      `json:"netbiosName"`
	Created          float64       `json:"created"`
	LastContact      float64       `json:"lastContact"`
	LastUpdate       float64       `json:"lastUpdate"`
	Tags             []string    `json:"tags"`
	Maintenance      Maintenance `json:"maintenance"`
	References       References  `json:"references"`
}

type Maintenance struct {
	Status        string `json:"status"`
	Start         float64  `json:"start"`
	End           float64  `json:"end"`
	ReasonMessage string `json:"reasonMessage"`
}

type References struct {
	Organization            Ation                   `json:"organization"`
	Location                Ation                   `json:"location"`
	RolePolicy              Policy                  `json:"rolePolicy"`
	Policy                  Policy                  `json:"policy"`
	Role                    Role                    `json:"role"`
	BackupUsage             map[string]int64        `json:"backupUsage"`
	Warranty                Warranty                `json:"warranty"`
	AssignedOwner           AssignedOwner           `json:"assignedOwner"`
}

type AssignedOwner struct {
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

type Ation struct {
	Name             string   `json:"name"`
	Address          *string  `json:"address,omitempty"`
	Description      string   `json:"description"`
	ID               int64    `json:"id"`
	NodeApprovalMode *string  `json:"nodeApprovalMode,omitempty"`
}

type Policy struct {
	ID               int64  `json:"id"`
	ParentPolicyID   int64  `json:"parentPolicyId"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	NodeClass        string `json:"nodeClass"`
	Updated          int64  `json:"updated"`
	NodeClassDefault bool   `json:"nodeClassDefault"`
}

type Role struct {
	ID               int64  `json:"id"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	NodeClass        string `json:"nodeClass"`
	Custom           bool   `json:"custom"`
	ChassisType      string `json:"chassisType"`
	Created          int64  `json:"created"`
	NodeRoleParentID int64  `json:"nodeRoleParentId"`
	Icon             string `json:"icon"`
}

type Warranty struct {
	StartDate                   float64 `json:"startDate"`
	EndDate                     float64 `json:"endDate"`
	ManufacturerFulfillmentDate float64 `json:"manufacturerFulfillmentDate"`
}

func (c *Client) GetDevices() ([]Device, error) {
	req, err := http.NewRequest(http.MethodGet, c.baseURL+"/v2/devices", nil)
	if err != nil {
		return nil, err
	}
	c.SetHeader(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	devices, err := UnmarshalDevices(body)
	if err != nil {
		return nil, err
	}

	return devices, nil
}
