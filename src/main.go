package main

import (
	"api/src/config"
	"api/src/controller"
	"api/src/middleware"
	"api/src/router"
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := flag.Int("port", 8000, "port to listen on")
	flag.Parse()

	cfg := config.New()
	if err := cfg.Init(int16(*port), ".env"); err != nil {
		panic(err)
	}
	ctl := controller.New(cfg)

	r := router.Gerar(ctl)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), middleware.RecoverPanic(middleware.EnableCors(r, cfg), cfg)))
}
