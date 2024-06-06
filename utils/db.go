package utils

import (
    "github.com/jmoiron/sqlx"
    _ "github.com/go-sql-driver/mysql"
    "log"
)

var BD *sqlx.DB

func InitBD() {
    var err error
    dsn := "root:@tcp(localhost)/tp-go?parseTime=true"
    BD, err = sqlx.Connect("mysql", dsn)
    if err != nil {
        log.Fatalln(err)
    }
}
