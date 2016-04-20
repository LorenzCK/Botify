package main

import (
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

type FastCGIServer struct {
}

func (s FastCGIServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
    resp.Write([]byte("<h1>Hello, 世界</h1>\n<p>Behold my Go web app.</p>"))
}

func main() {
    mode ServerMode = TCP
    for a := range os.Args[1:] {
        if(strings.EqualsFold(a, "-tcp")) {
            mode = TCP
        }
        else if(strings.EqualsFold(a, "-http")) {
            mode = HTTP
        }
    }

    /*
    listener, _ := net.Listen("tcp", "127.0.0.1:9999")
    srv := new(FastCGIServer)
    fcgi.Serve(listener, srv)*/

    fmt.Printf("Mode: %d\n", mode)
}
