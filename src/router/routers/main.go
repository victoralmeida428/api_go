package routers

import (
	"api/src/controller"
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	URI    string
	Metodo string
	Funcao func(http.ResponseWriter, *http.Request)
	Auth   bool
}

func Configurar(r *mux.Router, c *controller.Controller) *mux.Router {
	rotas := rotasHome
	for _, rota := range rotas(c) {
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}
	return r
}
