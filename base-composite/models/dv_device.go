package models

type DeviceRequest struct {
	Imei           string `json:"imei"`
	GpstypeID      int64  `json:"gpstype_id"`
	GSMNumber      string `json:"gsm_number"`
	Desription     string `json:"desription"`
	OrganizationID string `json:"organization_id"`
	CreateAt       string `json:"create_at"`
	UpdateAt       string `json:"update_at"`
	DeleteAt       string `json:"delete_at"`
}

type UpdateDeviceRequest struct {
	GpstypeID      int64  `json:"gpstype_id"`
	GSMNumber      string `json:"gsm_number"`
	Desription     string `json:"desription"`
	OrganizationID string `json:"organization_id"`
	CreateAt       string `json:"create_at"`
	UpdateAt       string `json:"update_at"`
	DeleteAt       string `json:"delete_at"`
}

type ListDeviceResult struct {
	NData     int            `json:"NData"`
	Page      int64          `json:"Page"`
	TotalPage int64          `json:"TotalPage"`
	TotalData int            `json:"TotalData"`
	From      int            `json:"From"`
	To        int            `json:"To"`
	Data      []DeviceResult `json:"Data"`
}

type DeviceResult struct {
	Imei           string `json:"imei"`
	GpstypeID      int64  `json:"gpstype_id"`
	GSMNumber      string `json:"gsm_number"`
	Desription     string `json:"desription"`
	OrganizationID string `json:"organization_id"`
	CreateAt       string `json:"create_at"`
	UpdateAt       string `json:"update_at"`
	DeleteAt       string `json:"delete_at"`
}
