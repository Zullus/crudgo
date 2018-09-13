package movierouter

import (
	"encoding/json"
	"net/http"

	. "github.com/Zullus/crudgo/config/dao"
	. "github.com/Zullus/crudgo/models"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

var dao = MoviesDAO{}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "applpication/json")
	w.WriteHeader(code)
	w.Write(response)
}

//GetAll retorna todos os filmes
func GetAll(w http.ResponseWriter, r *http.Request) {
	movies, err := dao.GetAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, movies)
}

//GetByID retorna o filme pelo ID
func GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie, err := dao.GetByID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Movie ID")
		return
	}
	respondWithJSON(w, http.StatusOK, movie)
}

//Create salva um novo filme
func Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var movie Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	movie.ID = bson.NewObjectId()
	if err := dao.Create(movie); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, movie)
}

//Update atualiza o filme
func Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	var movie Movie
	if err := dao.Update(params["id"], movie); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": movie.Name + " atualizado com sucesso!"})
}

//Delete apaga o filme
func Delete(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	if err := dao.Delete(params["id"]); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
