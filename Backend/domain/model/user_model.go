package model

import (
	"time"
)

type UserProject struct {
	Userid    int
	Firstname string
	Photouser string
}

type UserProjectDetail struct {
	Userid    int `json:"id"`
	Firstname string
	Lastname  string
	Photouser string
	Rolename  string
	Status    bool
}

type User struct {
	Userid         int    `gorm:"primary_key"`
	Email          string `gorm:"not null;"`
	Password       string `gorm:"not null; default:'';"`
	Status         bool   `gorm:"not null; default:false;"`
	Staterecovery  bool   `gorm:"not null; default: true;"`
	Mobileid       string `gorm:"not null; default:'';"`
	Roleid         int    `gorm:"not null;"`
	Resourcedataid int    `gorm:"not null;"`
	Rolename       string `sql:"-"`
	Resourcedata   Resourcedatas
}

type Resourcedatas struct {
	Resourcedataid int                `gorm:"primary_key" json:"resourcedataid"`
	Firstname      string             `gorm:"not null;" json:"firstname"`
	Lastname       string             `gorm:"not null;" json:"lastname"`
	Photouser      string             `gorm:"not null; default:''" json:"Image" `
	About          string             `gorm:"not null;" json:"about"`
	Cellphone      string             `gorm:"not null;" json:"cellphone"`
	Personalemail  string             `gorm:"not null;" json:"personalemail"`
	Datebirth      time.Time          `gorm:"not null;" json:"datebirth"`
	Address        string             `gorm:"not null;" json:"address"`
	Country        string             `gorm:"not null;" json:"country"`
	City           string             `gorm:"not null;" json:"city"`
	Cp             string             `gorm:"not null;" json:"cp"`
	Hardskills     []Hardsresource    `json:"hardskills"`
	Educations     []Education        `json:"educations"`
	Experiences    []Experience       `json:"experiences"`
	Softskills     []Softsresource    `json:"softskills"`
	Languages      []Languageresource `json:"languages"`
	Certificates   []Certificate      `json:"certificates"`
}

type MyData struct {
	Resourcedataid int                `gorm:"primary_key" json:"resourcedataid"`
	Rolename       string             `sql:"-" json:"rolename"`
	Firstname      string             `gorm:"not null;" json:"firstname"`
	Lastname       string             `gorm:"not null;" json:"lastname"`
	Photouser      string             `gorm:"not null; default:''" json:"Image" `
	About          string             `gorm:"not null;" json:"about"`
	Cellphone      string             `gorm:"not null;" json:"cellphone"`
	Personalemail  string             `gorm:"not null;" json:"personalemail"`
	Datebirth      time.Time          `gorm:"not null;" json:"datebirth"`
	Address        string             `gorm:"not null;" json:"address"`
	Country        string             `gorm:"not null;" json:"country"`
	City           string             `gorm:"not null;" json:"city"`
	Cp             string             `gorm:"not null;" json:"cp"`
	Hardskills     []Hardsresource    `json:"hardskills"`
	Educations     []Education        `json:"educations"`
	Experiences    []Experience       `json:"experiences"`
	Softskills     []Softsresource    `json:"softskills"`
	Languages      []Languageresource `json:"languages"`
	Certificates   []Certificate      `json:"certificates"`
}

type UserCard struct {
	Userid         int
	Roleid         int
	Rolename       string
	Firstname      string
	Lastname       string
	Photouser      string
	Status         bool
	Resourcedataid int
}

type GetUserDatas struct {
	Userid         int `json:"id"`
	Email          string
	Password       string
	Status         bool
	Staterecovery  bool
	Mobileid       string
	Roleid         int
	Resourcedataid int
	Rolename       string `sql:"-"`
	Resourcedata   Resourcedatas
}
