package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/jasonbronson/logreader/library/log"

	"github.com/jasonbronson/logreader/config"
)

func main() {

	newRelicApp := config.NewRelicApp()
	r := Router(newRelicApp)

	log.Printf("Listening to http://0.0.0.0:%v/", strconv.Itoa(config.Cfg.Port))

	port := "0.0.0.0:" + strconv.Itoa(config.Cfg.Port)
	srv := &http.Server{
		Addr: port,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 30,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of Gin in.
	}

	log.Fatal(srv.ListenAndServe())

}
