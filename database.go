package main

import (
    "fmt"
    "log"

    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

type Severity uint
const (
    DEBUG   Severity = 0
    INFO    Severity = 64
    WARNING Severity = 128
    ERROR   Severity = 255
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

func DbGetBotsForUser(userId int) (string, error) {
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

func DbLog(context *Context, severity Severity, tag string, message string) {
    conn, err := openBotifyDb()
    if(err != nil) {
        return
    }
    defer conn.Close()

    //TODO Add context here
    result, err := conn.Exec(
        "INSERT INTO `log` (timestamp, severity, tag, message) VALUES(NOW(), ?, ?, ?)",
        severity,
        tag,
        message,
    )
    if(err != nil) {
        log.Printf("Failed to write to log (%s)", err)
        return
    }
}
