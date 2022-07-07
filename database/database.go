package database

import (
	"fmt"
	"jwt-practice/entity"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func GetDatabase() *gorm.DB {

	//databasename := "userdb"
	database := "postgres"
	//databasepassword := "postgres"
	//"postgres://postgres:postgres@localhost:5432/userdb"
	databaseurl := "postgres://postgres:postgres@localhost:5432/userdb?sslmode=disable"
	//databaseurl := "postgres://postgres:postgres@localhost/userdb"
	fmt.Println("databaseurl:", databaseurl)

	connection, err := gorm.Open(database, databaseurl)
	if err != nil {
		log.Fatalln("wrong database url", err)
	}

	sqldb := connection.DB()

	err = sqldb.Ping()
	if err != nil {
		log.Fatal("database connected")
	}

	fmt.Println("connected to database")

	return connection

}
func InitialMigration() {
	connection := GetDatabase()
	defer Closedatabase(connection)
	connection.AutoMigrate(&entity.User{})
}
func Closedatabase(connection *gorm.DB) {
	sqldb := connection.DB()
	sqldb.Close()
}
