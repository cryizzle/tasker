package main

import (
	"fmt"
	"log"
	"os"

	"github.com/cryizzle/tasker/tasker_server/server"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	DATABASE_URL, DB_DRIVER, PORT string
)

func init() {
	mysql_user := os.Getenv("MYSQL_USER")
	mysql_password := os.Getenv("MYSQL_PASSWORD")
	mysql_database := os.Getenv("MYSQL_DATABASE")
	DB_DRIVER = os.Getenv("DB_DRIVER")
	PORT = os.Getenv("PORT")

	DATABASE_URL =
		fmt.Sprintf(
			"%s:%s@tcp(host.docker.internal:3306)/%s?parseTime=true", mysql_user, mysql_password, mysql_database,
		)

	log.Println("DATABASE_URL: ", DATABASE_URL)

}

func DBClient() (*sqlx.DB, error) {

	db, err := sqlx.Open(DB_DRIVER, DATABASE_URL)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	fmt.Println("Connected to DB")
	return db, nil
}

func main() {
	db, err := DBClient()
	if err != nil {
		log.Println(err)
		log.Fatalln("Couldn't connect to DB")
	}

	srv := server.CreateServer(db)
	srv.Routes()
	log.Println("server running on port:8000")
	srv.Start(PORT)
}
