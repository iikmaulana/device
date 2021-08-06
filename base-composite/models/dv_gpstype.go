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
	Name        string `json:"name"`
	Description string `json:"description"`
	CreateAt    string `json:"create_at"`
	UpdateAt    string `json:"update_at"`
	DeleteAt    string `json:"delete_at"`
}

type ListGpsTypeResult struct {
	NData     int             `json:"NData"`
	Page      int64           `json:"Page"`
	TotalPage int64           `json:"TotalPage"`
	TotalData int             `json:"TotalData"`
	From      int             `json:"From"`
	To        int             `json:"To"`
	Data      []GpsTypeResult `json:"Data"`
}

type GpsTypeResult struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreateAt    string `json:"create_at"`
	UpdateAt    string `json:"update_at"`
	DeleteAt    string `json:"delete_at"`
}
