package model

type Module struct {
	Moduleid   int `json:"id" gorm:"primary_key"`
	Modulename string
}
