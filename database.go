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
            DbConnectionBotifyName,
        )
    }

    log.Printf("Opening connection to %s\n", connectionStringBotify)

    return sql.Open("mysql", connectionStringBotify)
}

var connectionStringProgramo string = ""

func openProgramoDb() (*sql.DB, error) {
    if(connectionStringBotify == "") {
        connectionStringProgramo = fmt.Sprintf("%s:%s@%s/%s",
            DbConnectionUsername,
            DbConnectionPassword,
            DbConnectionProgramoHost,
            DbConnectionProgramoName,
        )
    }

    log.Printf("Opening connection to %s\n", connectionStringProgramo)

    return sql.Open("mysql", connectionStringProgramo)
}

func GetBotsForUser(userId int) (string, error) {
    conn, err := openBotifyDb()
    if(err != nil) {
        return "", err
    }
    defer conn.Close()

    rows, err := conn.Query("SELECT `token` FROM `bots` WHERE `admin_user_id` = ?", userId)
    if(err != nil) {
        return "", err
    }
    defer rows.Close()

    for rows.Next() {
        var token string
        err = rows.Scan(&token)
        if(err != nil) {
            return "", err
        }

        return token, nil
    }

    return "", nil
}
