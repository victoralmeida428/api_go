package controller

import (
	"api/src/utils"
	"net/http"
)

func (c *Controller) Home(w http.ResponseWriter, r *http.Request) {
	dados, err := c.repository.Grupo.FindAll()
	if err != nil {
		c.cfg.Error.ErrorReponse(w, r, http.StatusInternalServerError, err.Error())
	}
	if err := utils.WriteJSON(w, dados, http.StatusOK, nil); err != nil {
		c.cfg.Error.ErrorReponse(w, r, http.StatusInternalServerError, err.Error())
	}
}
