package routers

import (
	"api/src/controller"
	"net/http"
)

func rotasHome(ctl *controller.Controller) []Route {
	return []Route{{
		URI:    "/",
		Metodo: http.MethodGet,
		Auth:   false,
		Funcao: ctl.Home,
	},
	}
}
