package model

type Education struct {
	Educationid    int    `gorm:"primary_key" json:"educationid"`
	Academicname   string `gorm:"not null;" json:"academicname"`
	Academictitle  string `gorm:"not null;" json:"academictitle"`
	Year           int    `gorm:"not null;" json:"year"`
	Resourcedataid int    `gorm:"not null;" json:"-"`
}
