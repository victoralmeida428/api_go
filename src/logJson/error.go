package logJson

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Application struct {
	Logger *Logger
}

func writeJSON(w http.ResponseWriter, data interface{}, status int, headers http.Header) error {
	js, err := json.Marshal(data)

	if err != nil {
		return err
	}

	js = append(js, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

type envelope map[string]interface{}

func (app *Application) LogError(err error) {
	app.Logger.PrintError(err.Error(), nil)
}

func (app *Application) ErrorReponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	addrs := r.RequestURI

	env := envelope{"error": message, "uri": addrs}

	err := writeJSON(w, env, status, nil)
	if err != nil {
		app.LogError(err)
		w.WriteHeader(500)
	}
}

func (app *Application) ServerErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.LogError(err)

	message := "the server encoutered a problem and could not process your request"

	app.ErrorReponse(w, r, http.StatusInternalServerError, message)
}

func (app *Application) NotFountResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"

	app.ErrorReponse(w, r, http.StatusInternalServerError, message)
}

func (app *Application) MethodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	app.ErrorReponse(w, r, http.StatusInternalServerError, message)
}

func (app *Application) FailedValidationResponse(w http.ResponseWriter, r *http.Request, erros map[string]string) {
	app.ErrorReponse(w, r, http.StatusUnprocessableEntity, erros)
}

func (app *Application) EditConflictReponse(w http.ResponseWriter, r *http.Request) {
	message := "unable to update the record due to and edit conflict, please try again"
	app.ErrorReponse(w, r, http.StatusConflict, message)
}

func GetError(logger *Logger) *Application {
	return &Application{Logger: logger}
}
