package model

type Client struct {
	Clientid     int    `json:"id" gorm:"primary_key"`
	Clientname   string `gorm:"not null;"`
	Socialreason string `gorm:"not null;"`
	Photoclient  string `json:"Image" gorm:"not null; default:''"`
	Rfc          string `gorm:"not null;"`
	Address      string
	City         string
	Cp           string
	Cellphone    string
	Status       bool `gorm:"not null; default:true"`
	Countryid    int  `gorm:"not null;"`
}

type Clients struct {
	Clientid    int `json:"id"`
	Clientname  string
	Photoclient string `json:"Image"`
}

type ClientProyect struct {
	Cliente  Client
	Projects []*ProjectCard
}
