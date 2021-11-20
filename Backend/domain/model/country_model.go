package model

type Countrys struct {
	Countryid   int    `json:"id" gorm:"primary_key"`
	Countryname string `gorm:"not null;"`
}
