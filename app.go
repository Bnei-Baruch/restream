package main

import (
	"encoding/json"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run(port string) {
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Content-Length", "Accept-Encoding", "Content-Range", "Content-Disposition", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "DELETE", "POST", "PUT", "OPTIONS"})

	// Make empty json if db arg not found
	if _, err := ioutil.ReadFile(db_path); err != nil {
		str := "{}"
		if err = ioutil.WriteFile(db_path, []byte(str), 0644); err != nil {
			log.Fatal(err)
		}
	}

	http.ListenAndServe(port, handlers.CORS(originsOk, headersOk, methodsOk)(a.Router))
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/restream", a.getJsonDB).Methods("GET")
	a.Router.HandleFunc("/restream", a.saveJsonDB).Methods("PUT")
	a.Router.HandleFunc("/workflow/{ep}", a.putJson).Methods("PUT")
	a.Router.HandleFunc("/{ep}/status", a.statusJson).Methods("GET")
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"status": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
