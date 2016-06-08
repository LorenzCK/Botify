package main

import (
    "fmt"
    "log"

    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

var connectionStringBotify string = ""

func openBotifyDb() (*sql.DB, error) {
    if(connectionStringBotify == "") {
        connectionStringBotify = fmt.Sprintf("%s:%s@%s/%s",
            DbConnectionUsername,
            DbConnectionPassword,
            DbConnectionBotifyHost,
            DbConnectionBotifyName)
    }

    //log.Printf("Opening connection to %s\n", connectionStringBotify)

    return sql.Open("mysql", connectionStringBotify)
}

var connectionStringProgramo string = ""

func openProgramoDb() (*sql.DB, error) {
    if(connectionStringBotify == "") {
        connectionStringProgramo = fmt.Sprintf("%s:%s@%s/%s",
            DbConnectionUsername,
            DbConnectionPassword,
            DbConnectionProgramoHost,
            DbConnectionProgramoName)
    }

    //log.Printf("Opening connection to %s\n", connectionStringProgramo)

    return sql.Open("mysql", connectionStringProgramo)
}
