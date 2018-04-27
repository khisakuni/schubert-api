package cmd

import (
	"log"
	"net/http"

	"github.com/rs/cors"
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

	// Score routes
	r.Handle("/u/{userID}/scores", adapt(api.ListScores(), copyMgoSession(session))).Methods("GET")
	r.Handle("/u/{userID}/scores/{id}", adapt(api.GetScore(), copyMgoSession(session))).Methods("GET")
	r.Handle("/u/{userID}/scores", adapt(api.CreateScore(), copyMgoSession(session))).Methods("POST")
	r.Handle("/u/{userID}/scores/{id}", adapt(api.UpdateScore(), copyMgoSession(session))).Methods("PUT")

	// User routes
	r.Handle("/users", adapt(api.CreateUser(), copyMgoSession(session))).Methods("POST")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})

	handler := c.Handler(r)

	if err := http.ListenAndServe(":4000", handler); err != nil {
		log.Fatal(err)
	}
}
