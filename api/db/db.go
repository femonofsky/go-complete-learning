package db

import (
	"fmt"

	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

func Connect(
	host string,
	port string,
	dbName string,
	user string,
	password string,
) (db *xorm.Engine, err error) {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		host, port, user, dbName, password)

	db, err = xorm.NewEngine("postgres", connectionString)
	if err != nil {
		return
	}

	db.SetConnMaxLifetime(1000)
	_, err = db.DBMetas()
	if err != nil {
		return nil, err
	}

	return

}
