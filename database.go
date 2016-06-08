package main

import (
    "fmt"
    "log"

    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

var connectionStringBotify string = ""

func OpenBotifyDb() (*sql.DB, error) {
    if(connectionStringBotify == "") {
        connectionStringBotify = fmt.Sprintf("%s:%s@%s/%s",
            DbConnectionUsername,
            DbConnectionPassword,
            DbConnectionBotifyHost,
            DbConnectionBotifyName)
    }

    log.Printf("Opening connection to %s\n", connectionStringBotify)

    return sql.Open("mysql", connectionStringBotify)
}