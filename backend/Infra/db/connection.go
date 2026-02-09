package db

import (
	"ecommerce/config"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionstring(dbconfig *config.DBConfig) string {
	//user->postgres
	//password-> password
	//host-> localhost
	//port-> 5432
	//dbname-> ecommerce
	//sslmode-> disable
	connstring := fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s sslmode=%s",
		dbconfig.User,
		dbconfig.Password,
		dbconfig.Host,
		dbconfig.DB_Port,
		dbconfig.DBName,
		dbconfig.SSLMode,
	)
	return connstring

}

func NewConnection(dbconfig *config.DBConfig) (*sqlx.DB, error) {
	dbSource := GetConnectionstring(dbconfig)
	db, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		log.Println("Failed to connect to database:", err)
		return nil, err
	}
	return db, nil

}
