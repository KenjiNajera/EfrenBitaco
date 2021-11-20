package model

type Permission struct {
	Permissionid   int `json:"id" gorm:"primary_key"`
	Permissionname string
}
