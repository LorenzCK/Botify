package main

import (
    "log"
    "fmt"

    "net/http"
    "github.com/julienschmidt/httprouter"
)

func TestIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    log.Printf("Access to '%s' from %s\n", r.RequestURI, r.RemoteAddr)

    w.Write([]byte("<h1>Hello, 世界</h1>\n<p>Behold my Go web app.</p>"))

    result, err := DbGetBotsForUser(1)
    if(err != nil) {
        w.Write([]byte(fmt.Sprintf("<p>Error: %s</p>", err)))
    } else {
        w.Write([]byte(fmt.Sprintf("<p>Bot token for user 1: %s</p>", result)))
    }
}

func RouteBotifyHook(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    log.Printf("Access to '%s' from %s\n", r.RequestURI, r.RemoteAddr)

    w.Write([]byte("<h1>Botify hook</h1>\n<p>Behold my Go web app.</p>"))
}

func RouteHostedHook(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    token := ps.ByName("token")

    log.Printf("Access to '%s' from %s (token %s)\n", r.RequestURI, r.RemoteAddr, token)

    //Noop
}

func SetupRouter() *httprouter.Router {
    router := httprouter.New()
    router.GET("/", TestIndex)
    router.POST(fmt.Sprintf("/hook/%s", BotifyBotToken), RouteBotifyHook)
    router.POST("/bot/hook/:token", RouteHostedHook)

    return router
}
