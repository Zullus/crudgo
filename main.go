package main

import (
	"fmt"
	"log"
	"net/http"

	. "github.com/Zullus/crudgo/config"
	. "github.com/Zullus/crudgo/config/dao"
	"github.com/gorilla/mux"

	moviesrouter "github.com/Zullus/crudgo/router"
)

var dao = MoviesDAO{}
var config = Config{}

func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/movies", moviesrouter.GetAll).Methods("GET")
	r.HandleFunc("/api/v1/movies/{id}", moviesrouter.GetByID).Methods("GET")
	r.HandleFunc("/api/v1/movies", moviesrouter.Create).Methods("POST")
	r.HandleFunc("/api/v1/movies/{id}", moviesrouter.Update).Methods("PUT")
	r.HandleFunc("/api/v1/movies/{id}", moviesrouter.Delete).Methods("DELETE")

	var port = ":3000"
	fmt.Println("Server running in port: ", port)
	log.Fatal(http.ListenAndServe(port, r))
}
