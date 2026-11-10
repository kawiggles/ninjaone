package ninja

import (
	"encoding/json"
	"io"
	"net/http"
)

type Devices []Device

func UnmarshalDevices(data []byte) (Devices, error) {
	var r Devices
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Devices) Marshal() ([]byte, error) {
	return json.Marshal(r)
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
	Created          int64       `json:"created"`
	LastContact      int64       `json:"lastContact"`
	LastUpdate       int64       `json:"lastUpdate"`
	UserData         UserData    `json:"userData"`
	Tags             []string    `json:"tags"`
	Maintenance      Maintenance `json:"maintenance"`
	References       References  `json:"references"`
}

type Maintenance struct {
	Status        string `json:"status"`
	Start         int64  `json:"start"`
	End           int64  `json:"end"`
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
	BackupBandwidthThrottle BackupBandwidthThrottle `json:"backupBandwidthThrottle"`
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

type BackupBandwidthThrottle struct {
	Enabled              bool         `json:"enabled"`
	WorkHoursKbps        int64        `json:"workHoursKbps"`
	NonWorkHoursKbps     int64        `json:"nonWorkHoursKbps"`
	WorkHoursUserUnit    string       `json:"workHoursUserUnit"`
	NonWorkHoursUserUnit string       `json:"nonWorkHoursUserUnit"`
	WorkSchedule         WorkSchedule `json:"workSchedule"`
}

type WorkSchedule struct {
	EndHour     int64    `json:"endHour"`
	EndMinute   int64    `json:"endMinute"`
	StartHour   int64    `json:"startHour"`
	StartMinute int64    `json:"startMinute"`
	WeekDays    []string `json:"weekDays"`
}

type Ation struct {
	Name             string   `json:"name"`
	Address          *string  `json:"address,omitempty"`
	Description      string   `json:"description"`
	UserData         UserData `json:"userData"`
	ID               int64    `json:"id"`
	NodeApprovalMode *string  `json:"nodeApprovalMode,omitempty"`
}

type UserData struct {
	AdditionalProp1 AdditionalProp1 `json:"additionalProp1"`
}

type AdditionalProp1 struct {
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
	StartDate                   int64 `json:"startDate"`
	EndDate                     int64 `json:"endDate"`
	ManufacturerFulfillmentDate int64 `json:"manufacturerFulfillmentDate"`
}

func (c *Client) GetDevices() ([]Device, error) {
	resp, err := http.Get(c.baseURL+"/v2/devices")
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
