package db

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func GetConnectionstring() string {
	//user->postgres
	//password-> password
	//host-> localhost
	//port-> 5432
	//dbname-> ecommerce
	//sslmode-> disable
	return "user=postgres password=password host=localhost port=5432 dbname=ecommerce sslmode=disable"

}

func NewConnection() (*sqlx.DB, error) {
	dbSource := GetConnectionstring()
	db, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		log.Println("Failed to connect to database:", err)
		return nil, err
	}
	return db, nil

}
