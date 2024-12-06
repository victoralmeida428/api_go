package main

import (
	"api/src/config"
	"api/src/controller"
	"api/src/middleware"
	"api/src/model"
	"api/src/repository"
	"api/src/router"
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := flag.Int("port", 8000, "port to listen on")
	env := flag.String("env", ".env", "path of the file .env")
	flag.Parse()

	cfg := config.New()
	if err := cfg.Init(int16(*port), *env); err != nil {
		panic(err)
	}

	md := model.New(cfg.DB)
	repo := repository.New(md)
	ctl := controller.New(cfg, md, repo)

	r := router.Gerar(ctl)
	fmt.Printf("API running at port %d", cfg.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), middleware.RecoverPanic(middleware.EnableCors(r, cfg), cfg)))
}
