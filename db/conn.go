package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "D0uGu57$"
	dbname   = "mkt_db"
)

func ConnectDB() (*sql.DB, error) {

	//define a connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	//abre a conexão com o postgres
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	//executa um ping para checar se a conexão responde
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to" + dbname)

	return db, nil
}
