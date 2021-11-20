package model

type Typerequest struct {
	Typerequestid   int    `json:"id" gorm:"primary_key"`
	Typerequestname string `json:"typerequestname"`
}
