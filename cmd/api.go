package cmd

import (
	"log"
	"net/http"

	mgo "gopkg.in/mgo.v2"

	"github.com/gorilla/mux"
	"github.com/khisakuni/schubert-api/api"
	"github.com/khisakuni/schubert-api/api/middleware"
)

// RunAPI starts the server.
func RunAPI() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	adapt := middleware.Adapt
	copyMgoSession := middleware.CopyMgoSession
	r := mux.NewRouter()
	r.Handle("/scores", adapt(api.ListScores(), copyMgoSession(session))).Methods("GET")
	r.Handle("/scores/{id}", adapt(api.GetScore(), copyMgoSession(session))).Methods("GET")
	r.Handle("/scores", adapt(api.CreateScore(), copyMgoSession(session))).Methods("POST")
	r.Handle("/scores/{id}", adapt(api.UpdateScore(), copyMgoSession(session))).Methods("PUT")

	if err := http.ListenAndServe(":4000", r); err != nil {
		log.Fatal(err)
	}
}
