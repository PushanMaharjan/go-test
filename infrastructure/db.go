package infrastructure

import (
	"fmt"
	"go-fx-test/lib"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// db model
type Database struct {
	*gorm.DB
	dsn string
}

// new database creates new instance
func NewDatabase(env lib.Env) Database {
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", env.DBUsername, env.DBPassword, env.DBHost, env.DBPort, env.DBName)

	db, err := gorm.Open(mysql.Open(url))

	if err != nil {
		log.Println("Url: ", url)
		log.Panic(err)
	}

	log.Println("database connected")

	return Database{
		DB:  db,
		dsn: url,
	}
}
