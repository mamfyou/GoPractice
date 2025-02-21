package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	dbUser = "postgres"
	dbPass = "Mamf20021381@"
	dbHost = "localhost"
	dbPort = 5432
)

var db *gorm.DB

func init(){
	session := postgres.Open(fmt.Sprintf("host %s", dbHost))
}

func GetDB() *gorm.DB {
	return db
}