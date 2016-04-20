package main

import (
    "fmt"
    "strings"

    "net"
    "net/http"
//    "net/http/fcgi"

    "os"

//    "github.com/julienschmidt/httprouter"

//    "database/sql"
//    "github.com/go-sql-driver/mysql"
)

type ServerMode int

const (
    HTTP ServerMode = iota
    TCP
)

func main() {
    // Command line parsing
    var mode = TCP
    for _, a := range os.Args[1:] {
        if(strings.EqualFold(a, "-tcp")) {
            mode = TCP
        } else if(strings.EqualFold(a, "-http")) {
            mode = HTTP
        }
    }

}
