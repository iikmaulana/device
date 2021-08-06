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
	GpstypeID      int64  `db:"gpstype_id" json:"gpstype_id"`
	GSMNumber      string `db:"gsm_number" json:"gsm_number"`
	Desription     string `db:"desription" json:"desription"`
	OrganizationID string `db:"organization_id" json:"organization_id"`
	CreateAt       string `db:"create_at" json:"create_at"`
	UpdateAt       string `db:"update_at" json:"update_at"`
	DeleteAt       string `db:"delete_at" json:"delete_at"`
}

type ListDeviceResult struct {
	Result FormMetaData   `json:"result,omitempty"`
	Data   []DeviceResult `json:"data,omitempty"`
}

type DeviceResult struct {
	Imei           string `db:"imei" json:"imei"`
	GpstypeID      int64  `db:"gpstype_id" json:"gpstype_id"`
	GSMNumber      string `db:"gsm_number" json:"gsm_number"`
	Desription     string `db:"desription" json:"desription"`
	OrganizationID string `db:"organization_id" json:"organization_id"`
	CreateAt       string `db:"create_at" json:"create_at"`
	UpdateAt       string `db:"update_at" json:"update_at"`
	DeleteAt       string `db:"delete_at" json:"delete_at"`
}
