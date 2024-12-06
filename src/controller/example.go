package controller

import (
	"api/src/utils"
	"net/http"
)

func (c *Controller) Home(w http.ResponseWriter, r *http.Request) {
	var status = http.StatusOK
	defer c.cfg.Logger.PrintDebug(r, status, nil)

	dados, err := c.repository.Grupo.FindAll()

	if err != nil {
		status = http.StatusInternalServerError
		c.cfg.Error.ErrorReponse(w, r, status, err.Error())
		return
	}
	if err := utils.WriteJSON(w, dados, status, nil); err != nil {
		status = http.StatusInternalServerError
		c.cfg.Error.ErrorReponse(w, r, status, err.Error())
		return
	}
}
