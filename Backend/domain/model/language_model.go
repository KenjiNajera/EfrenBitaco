package model

type Language struct {
	Languageid   int    `json:"id" gorm:"primary_key" example:"0"`
	Languagename string `gorm:"not null;" example:"Español"`
}

type Languageresource struct {
	Languageresourceid int    `gorm:"primary_key" json:"languageresourceid"`
	Domain             int    `gomr:"not null; default:0" json:"domain"`
	Languageid         int    `gomr:"not null;" json:"languageid"`
	Resourcedataid     int    `gomr:"not null;" json:"-"`
	Languagename       string `sql:"-" json:"languagename"`
}
