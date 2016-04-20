package main

import (
    "log"
    "fmt"
    "strings"
    "net"

    "net/http"
    "net/http/fcgi"

    "os"

    "github.com/julienschmidt/httprouter"

//    "database/sql"
//    "github.com/go-sql-driver/mysql"
)

type ServerMode int

const (
    HTTP ServerMode = iota
    TCP
)

func TestIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    log.Printf("Access to / from %s\n", r.RemoteAddr)

    w.Write([]byte("<h1>Hello, 世界</h1>\n<p>Behold my Go web app.</p>"))
}

func SetupRouter() *httprouter.Router {
    router := httprouter.New()
    router.GET("/", TestIndex)

    return router
}

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
            logfile := fmt.Sprintf("%s.log", os.Args[0])
            log.Printf("Using file %s for log output", logfile)
            //TODO use actual logfile!

            f, err := os.OpenFile("server.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
            if err != nil {
                log.Fatal("Failed to open log file")
            }
            log.SetOutput(f)
        }
    }

    log.Println("Setting up router")
    router:= SetupRouter()

    if(mode == TCP) {
        log.Println("Starting to serve on TCP port 9999")
        listener, _ := net.Listen("tcp", "127.0.0.1:9999")
        log.Fatal(fcgi.Serve(listener, router))
    } else if(mode == HTTP) {
        log.Println("Starting to serve on HTTP port 8080")
        log.Fatal(http.ListenAndServe(":8080", router))
    } else {
        log.Fatal("No suitable listening protocol selected")
    }
}
