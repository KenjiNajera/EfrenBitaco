package model

import "time"

type Request struct {
	Requestid     int
	Description   string
	Status        bool
	Createat      time.Time
	Details       string
	Userid        int
	Typerequestid int
}

type GetRequest struct {
	Requestid       int `json:"id"`
	Firstname       string
	Lastname        string
	Photouser       string
	Typerequestname string
	Createat        time.Time `json:"-"`
	CreateAt        string
	Description     string
}

func (gtr *GetRequest) Format() string {
	gtr.CreateAt = gtr.Createat.Format("2006-01-02")
	return gtr.CreateAt
}

type GetRequestUser struct {
	Requestid       int    `json:"requestid"`
	Status          *bool  `json:"status"`
	Typerequestname string `json:"typerequestname"`
	Description     string `json:"description"`
	Details         string `json:"details"`
}
