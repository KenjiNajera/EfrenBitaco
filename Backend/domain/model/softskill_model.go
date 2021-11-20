package model

type Softskill struct {
	Softskillid   int    `json:"id" gorm:"primary_key"`
	Softskillname string `gorm:"not null;"`
}

type Softsresource struct {
	Softsresourceid int    `gorm:"primary_key" json:"softsresourceid"`
	Softskillid     int    `gorm:"not null;" json:"softskillid"`
	Resourcedataid  int    `gorm:"not null:" json:"-"`
	Softskillname   string `sql:"-"  json:"softskillname"`
}
