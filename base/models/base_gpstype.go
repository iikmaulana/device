package models

type GpsTypeRequest struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreateAt    string `json:"create_at"`
	UpdateAt    string `json:"update_at"`
	DeleteAt    string `json:"delete_at"`
}

type UpdateGpsTypeRequest struct {
	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
	CreateAt    string `db:"create_at" json:"create_at"`
	UpdateAt    string `db:"update_at" json:"update_at"`
	DeleteAt    string `db:"delete_at" json:"delete_at"`
}

type ListGpsTypeResult struct {
	Result FormMetaData    `json:"result,omitempty"`
	Data   []GpsTypeResult `json:"data,omitempty"`
}

type GpsTypeResult struct {
	ID          int64  `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
	CreateAt    string `db:"create_at" json:"create_at"`
	UpdateAt    string `db:"update_at" json:"update_at"`
	DeleteAt    string `db:"delete_at" json:"delete_at"`
}
