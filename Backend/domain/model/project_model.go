package model

type Project struct {
	Projectid   int    `gorm:"primary_key"`
	Projectname string `gorm:"not null;"`
	Description string `gorm:"not null; default:''"`
	Status      bool   `gorm:"not null; default:false"`
	Clientid    int    `gorm:"not null;"`
	Hardskills  []int  `gorm:"-"`
}

type ProjectCard struct {
	Projectid   int `json:"id"`
	Projectname string
	Description string
	Status      bool
	Clientid    int
	Hardskills  []*Hardskill
	Resources   []*UserProject
}

type ProjectCardDetail struct {
	Projectid   int `json:"id"`
	Projectname string
	Description string
	Status      bool
	Clientid    int
	Hardskills  []*Hardskill
	Resources   []*UserProjectDetail
}
