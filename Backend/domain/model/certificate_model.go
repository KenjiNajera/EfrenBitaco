package model

type Certificate struct {
	Certificateid   int    `gorm:"primary_key" json:"certificateid"`
	Idcertificate   string `gorm:"not null;" json:"idcertificate"`
	Certificatename string `gorm:"not null;" json:"certificatename"`
	Expedition      string `gorm:"not null;" json:"expedition"`
	Expiration      string `gorm:"not null; default:'(without caducity)';" json:"expiration"`
	Resourcedataid  int    `gorm:"not null;" json:"-"`
}
