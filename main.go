package main

import (
	"log"
	"net/http"
	"time"

	"github.com/juliotorresmoreno/fidelapp-realtime/router"
	"github.com/juliotorresmoreno/fidelapp-realtime/ws"
)

func main() {
	hub := ws.NewHub()
	router := router.NewRouter()

	srv := &http.Server{
		Handler: router,
		Addr:    ":4000",

		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ws.ServeWs(hub, w, r)
	})

	log.Println("Listening at " + srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
