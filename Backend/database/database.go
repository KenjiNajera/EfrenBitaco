package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

const (
	host     = "ec2-52-203-182-92.compute-1.amazonaws.com"
	port     = "5432"
	user     = "yjdbasvacfxrkj"
	password = "c03ef4055284bf480f665319f571768485c8b5068b3053e7dcbb8be2cba28c4e"
	dbname   = "d8f4itgc343emh"
)

func InitDB() *gorm.DB {

	dsn := "postgres://" + user + ":" + password + "@" + host + ":" + port + "/" + dbname
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		panic(err.Error())
	}

	return db
}
