package main

import (
    "log"
    "net/http"
    "github.com/julienschmidt/httprouter"
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
