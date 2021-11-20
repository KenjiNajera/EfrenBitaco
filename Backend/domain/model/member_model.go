package model

type Member struct {
	Memberid  int
	Userid    int
	Projectid int
	Leader    bool `json:"-" gorm:"-"`
}
