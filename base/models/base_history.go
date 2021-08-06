package models

type HistoryRequest struct {
	ID             string `json:"id"`
	Imei           string `json:"imei"`
	VehicleID      string `json:"vehicle_id"`
	Description    string `json:"description"`
	Time           string `json:"time"`
	Status         int64  `json:"status"`
	OrganizationID string `json:"organization_id"`
	CreateAt       string `json:"create_at"`
	UpdateAt       string `json:"update_at"`
	DeleteAt       string `json:"delete_at"`
}

type UpdateHistoryRequest struct {
	Imei           string `db:"imei" json:"imei"`
	VehicleID      string `db:"vehicle_id" json:"vehicle_id"`
	Description    string `db:"description" json:"description"`
	Time           string `db:"time" json:"time"`
	Status         int64  `db:"status" json:"status"`
	OrganizationID string `db:"organization_id" json:"organization_id"`
	CreateAt       string `db:"create_at" json:"create_at"`
	UpdateAt       string `db:"update_at" json:"update_at"`
	DeleteAt       string `db:"delete_at" json:"delete_at"`
}

type ListHistoryResult struct {
	Result FormMetaData    `json:"result,omitempty"`
	Data   []HistoryResult `json:"data,omitempty"`
}

type HistoryResult struct {
	ID             string `db:"id" json:"id"`
	Imei           string `db:"imei" json:"imei"`
	VehicleID      string `db:"vehicle_id" json:"vehicle_id"`
	Description    string `db:"description" json:"description"`
	Time           string `db:"time" json:"time"`
	Status         int64  `db:"status" json:"status"`
	OrganizationID string `db:"organization_id" json:"organization_id"`
	CreateAt       string `db:"create_at" json:"create_at"`
	UpdateAt       string `db:"update_at" json:"update_at"`
	DeleteAt       string `db:"delete_at" json:"delete_at"`
}
