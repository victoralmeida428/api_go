package router

import (
	"api/src/controller"
	"api/src/router/routers"
	"github.com/gorilla/mux"
)

func Gerar(c *controller.Controller) *mux.Router {
	r := mux.NewRouter()
	return routers.Configurar(r, c)
}
