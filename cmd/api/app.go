package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type (
	App struct {
		Router *mux.Router
	}
)

func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.initRoutes()
}

func (a *App) Run(addr string) {
	log.Printf("Listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) initRoutes() {
	a.Router.HandleFunc("/version", a.getVersion).Methods("GET")
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, "{\"error\":\"failed to marshal response\"}", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}
