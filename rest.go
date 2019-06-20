package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"os"
)

func (a *App) getJsonDB(w http.ResponseWriter, r *http.Request) {
	var v map[string]interface{}

	mu.Lock()
	d, err := ioutil.ReadFile(db_path)
	mu.Unlock()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	err = json.Unmarshal(d, &v)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, v)
}

func (a *App) saveJsonDB(w http.ResponseWriter, r *http.Request) {
	var j map[string]interface{}

	d := json.NewDecoder(r.Body)
	if err := d.Decode(&j); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}

	b, err := json.Marshal(j)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	mu.Lock()
	f, err := os.Create(db_path)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		mu.Unlock()
		return
	}
	_, err = f.Write(b)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		mu.Unlock()
		return
	}
	mu.Unlock()

	respondWithJSON(w, http.StatusOK, map[string]string{"status": "success"})
}

func (a *App) putJson(w http.ResponseWriter, r *http.Request) {
	var s status
	vars := mux.Vars(r)
	endpoint := vars["ep"]

	b, _ := ioutil.ReadAll(r.Body)

	err := s.putExec(endpoint, string(b))

	defer r.Body.Close()

	if err != nil {
		s.Status = "error"
	} else {
		s.Status = "ok"
	}

	respondWithJSON(w, http.StatusOK, s)
}

func (a *App) statusJson(w http.ResponseWriter, r *http.Request) {
	var s status
	vars := mux.Vars(r)
	endpoint := vars["ep"]
	id := r.FormValue("id")
	key := r.FormValue("key")
	value := r.FormValue("value")

	err := s.getStatus(endpoint, id, key, value)

	if err != nil {
		s.Status = "error"
	} else {
		s.Status = "ok"
	}

	respondWithJSON(w, http.StatusOK, s)
}