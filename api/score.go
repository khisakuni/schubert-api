package api

import (
	"fmt"
	"net/http"

	"github.com/khisakuni/schubert-api/api/middleware"
)

func ListScores(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Scores List")
}

func GetScore() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session := middleware.DBFromContext(r.Context())
		fmt.Printf("Session from context >> %v", session)
	})
}

func CreateScore(w http.ResponseWriter, r *http.Request) {}

func UpdateScore(w http.ResponseWriter, r *http.Request) {}
