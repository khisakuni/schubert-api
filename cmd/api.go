package cmd

import (
	"log"
	"net/http"

	"github.com/globalsign/mgo"

	"github.com/gorilla/mux"
	"github.com/khisakuni/schubert-api/api"
	"github.com/khisakuni/schubert-api/api/middleware"
)

func RunAPI() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	r := mux.NewRouter()
	r.Handle("/scores/{id}", middleware.Adapt(api.GetScore(), middleware.CopyMgoSession(session))).Methods("GET")
	r.HandleFunc("/scores", api.CreateScore).Methods("POST")
	r.HandleFunc("/scores/{id}", api.UpdateScore).Methods("PUT")

	if err := http.ListenAndServe(":4000", r); err != nil {
		log.Fatal(err)
	}
}
