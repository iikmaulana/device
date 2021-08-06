package models

type FormMetaData struct {
	NData     int
	Page      int
	TotalPage int64
	TotalData int64
	From      int
	To        int
	Data      interface{}
}
