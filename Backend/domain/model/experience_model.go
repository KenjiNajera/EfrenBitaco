package model

type Experience struct {
	Experienceid    int    `gorm:"primary_key" json:"experienceid"`
	Projecttitle    string `gorm:"not null;" json:"projecttitle"`
	Description     string `gorm:"not null;" json:"description"`
	Yearelaboration int    `gorm:"not null;" json:"yearelaboration"`
	Resourcedataid  int    `json:"-"`
}
