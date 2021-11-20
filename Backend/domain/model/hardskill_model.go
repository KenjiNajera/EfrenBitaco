package model

type Hardskill struct {
	Hardskillid   int    `json:"id" form:"id" gorm:"primary_key"`
	Hardskillname string `gorm:"not null;"`
	Image         string `gorm:"not null;"`
}

type Hardsresource struct {
	Hardsresourceid int    `gorm:"primary_key" json:"hardsresourceid"`
	Domain          int    `gorm:"not null; default: 0;" json:"domain"`
	Resourcedataid  int    `gorm:"not null;" json:"-"`
	Hardskillid     int    `gorm:"not null;" json:"hardskillid"`
	Hardskillname   string `sql:"-" json:"hardskillname"`
	Image           string `sql:"-" json:"image"`
}
