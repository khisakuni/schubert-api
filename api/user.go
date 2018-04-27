package api

import (
	"encoding/json"
	"net/http"

	"github.com/khisakuni/schubert-api/api/middleware"

	"gopkg.in/mgo.v2/bson"
)

// User is user model.
type User struct {
	ID    bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Email string        `json:"email"`
}

// CreateUser endpoint creates user given an email.
func CreateUser() http.Handler {
	return appHandler(func(w http.ResponseWriter, r *http.Request) *apiError {
		var u User
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			return internalServerError(err)
		}

		// TODO: Validate presence of valid email.

		u.ID = bson.NewObjectId()
		session := middleware.MgoSessionFromContext(r.Context())
		err = session.DB("schubert").C("users").Insert(u)
		if err != nil {
			return internalServerError(err)
		}

		js, err := json.Marshal(u)
		if err != nil {
			return internalServerError(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)
		w.Write(js)
		return nil
	})
}
