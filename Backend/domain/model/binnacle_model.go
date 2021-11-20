package model

import (
	"time"
)

type Binnacle struct {
	Binnacleid  int `json:id`
	Monthname   string
	Month       int
	Year        int
	Totalhours  time.Time
	Createat    time.Time
	Status      bool
	Projectname string
	Userid      int
}

type CreateBinnacle struct {
	Month  int
	Year   int
	Userid int
	Status bool `json:"-"`
}
type FisnisBinnacle struct {
	Binnacleid int
	Status     bool
	Hourstotal float64
}

type BinnacleView struct {
	BinnacleData BinnacleViewData   `json:"binnacledata"`
	Projectnames []string           `json:"projectnames"`
	Days         []*DayBinnacleView `json:"days"`
}
type BinnacleViewUser struct {
	Binnacleid   int                    `json:"binnacleid"`
	Monthname    string                 `json:"monthname"`
	Year         int                    `json:"year"`
	Totalhours   float64                `json:"totalhours"`
	Status       bool                   `json:"status"`
	Projectnames []string               `json:"projectnames"`
	Days         []*DayBinnacleViewUser `json:"days"`
}

type DayBinnacleView struct {
	Daydata       *Day             `json:"daydata"`
	Checktime     ChecktimeDay     `json:"checktime"`
	Mealtime      Mealtime         `json:"mealtime"`
	Departuretime DeparturetimeDay `json:"departuretime"`
	Overtime      Overtime         `json:"overtime"`
	Activity      ActivityDay      `json:"activity"`
}
type DayBinnacleViewUser struct {
	Dayid         int       `json:"dayid"`
	Dayname       string    `json:"dayname"`
	Daydate       time.Time `json:"daydate"`
	Totalhoursday time.Time `json:"totalhoursday"`
	Checktime     time.Time `json:"checktime"`
	Departuretime time.Time `json:"departuretime"`
	Overtime      time.Time `json:"overtime"`
	Mealtime      time.Time `json:"mealtime"`
	Summary       string    `json:"summary"`
	Invoice       string    `json:"invoice"`
}

type BinnacleViewData struct {
	Binnacleid int
	Monthname  string
	Year       int
	Totalhours float64
	Status     bool
	Userid     int
	Firstname  string
	Lastname   string
	Photouser  string
}

type CreateDay struct {
	Dayid      int
	Binnacleid int
	Month      int
	Year       int
	Projectids []int
}

type Departuretime struct {
	Departuretime time.Time `json:"departuretime"`
	Daydate       time.Time `json:"daydate"`
	Overtime      bool      `json:"overtime"`
	Summary       string    `json:"summary"`
	Invoice       string    `json:"invoice"`
	Userid        int       `json:"userid"`
}
type CreateDeparureTime struct {
	Binnacleid int
	Dayid      int
	Checktime  int
	Mealtime   int
}

type Day struct {
	Dayid         int
	Dayname       string
	Daydate       time.Time
	Createat      time.Time
	Totalhoursday time.Time
	Dailytime     time.Time
}
type ChecktimeDay struct {
	Checktimeid int
	Checktime   time.Time
	Createat    time.Time
	Dayid       int
}
type DeparturetimeDay struct {
	Departuretimeid int
	Departuretime   time.Time
	Createat        time.Time
}
type ActivityDay struct {
	Activityid int
	Summary    string
	invoice    string
}

type Mealtime struct {
	Mealtimeid int
	Mealtime   time.Time `json:"mealtime"`
	Createat   time.Time
	Dayid      int
	Daydate    time.Time `json:"daydate"`
	Userid     int       `json:"userid"`
}
type CreateMealTime struct {
	Binnacleid      int
	Dayid           int
	DeparturetimeId int
}

type Checktime struct {
	Checktimeid int
	Checktime   time.Time
	Createat    time.Time
	Dayid       int
	Userid      int
	Binnacleid  int
	Daydate     time.Time
}

type Overtime struct {
	Overtimeid int
	Overtime   time.Time
	Createat   time.Time
	Dayid      int
}

type Activity struct {
	Activityid int
	Summary    string
	Invoice    string
	Dayid      int
}

type UpdateDay struct {
	Dayid         int
	Daydate       time.Time
	Totalhoursday time.Time
	Dailytime     time.Time
	Checktime     time.Time
	Mealtime      time.Time
	Departuretime time.Time
	Overtime      time.Time
	Summary       string
	Invoice       string
}
