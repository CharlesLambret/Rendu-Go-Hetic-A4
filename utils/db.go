package utils

import (
    "github.com/jmoiron/sqlx"
    _ "github.com/lib/pq"
    "log"
)

var BD *sqlx.DB

func InitBD() {
    var err error
    BD, err = sqlx.Connect("postgres", "user=youruser dbname=yourdb sslmode=disable")
    if err != nil {
        log.Fatalln(err)
    }
}
