package main

import (
    "fmt"
    "strings"

    "path"
    "os"
    "log"

    "net"
    "net/http"
    "net/http/fcgi"

//    "database/sql"
//    "github.com/go-sql-driver/mysql"
)

type ServerMode int

const (
    HTTP ServerMode = iota
    TCP
)

func main() {
    log.Println("Starting up")

    // Command line parsing
    var mode = TCP
    for _, a := range os.Args[1:] {
        if(strings.EqualFold(a, "-tcp")) {
            mode = TCP
        } else if(strings.EqualFold(a, "-http")) {
            mode = HTTP
        } else if(strings.EqualFold(a, "-logFile")) {
            exeName := path.Base(os.Args[0])
            logfile := fmt.Sprintf("%s%s", exeName, LogExtension)
            log.Printf("Using file %s for log output\n", logfile)

            f, err := os.OpenFile(logfile, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
            if err != nil {
                log.Fatalln("Failed to open log file")
            }
            log.SetOutput(f)
        }
    }

    log.Println("Setting up router")
    router:= SetupRouter()

    if(mode == TCP) {
        log.Printf("Starting to serve on TCP %s\n", TCPListenerAddress)
        listener, _ := net.Listen("tcp", TCPListenerAddress)
        log.Fatalln(fcgi.Serve(listener, router))
    } else if(mode == HTTP) {
        log.Printf("Starting to serve on HTTP %s\n", HTTPListenerAddress)
        log.Fatalln(http.ListenAndServe(HTTPListenerAddress, router))
    } else {
        log.Fatalln("No suitable listening protocol selected")
    }
}
