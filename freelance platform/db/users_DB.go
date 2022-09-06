package db

import (
	"database/sql"
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	login    = "postgres"
	password = "123"
	dbname   = "postgres"
)

type Pool struct {
	ID       int
	Login    string
	Password string
	Access   string
}

func Conectdb() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, login, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	var myUser Pool
	userSql := "SELECT id, login, password , access FROM \"useres\" "
	err = db.QueryRow(userSql, 1).Scan(&myUser.ID, &myUser.Login, &myUser.Password, &myUser.Access)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hi %s, welcome back!\n", myUser.Login)

}

func Registration(arga, argb, argc string) bool {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, login, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	if len(arga) <= 10 && len(argb) <= 10 && len(argc) <= 2 {
		_, err := db.Exec("insert into \"useres\" (login, password, access) values ($1, $2, $3)",
			arga, argb, argc)
		if err != nil {
			panic(err)
		}

	} else {
		return false
	}
	return true

}
