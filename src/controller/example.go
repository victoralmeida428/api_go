package controller

import (
	"api/src/utils"
	"net/http"
)

func (c *Controller) Home(w http.ResponseWriter, r *http.Request) {

	if err := utils.WriteJSON(w, c.cfg, http.StatusOK, nil); err != nil {
		c.cfg.Error.ErrorReponse(w, r, http.StatusInternalServerError, err.Error())
	}
}
