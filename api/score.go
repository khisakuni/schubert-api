package api

import (
	"encoding/json"
	"net/http"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
	"github.com/khisakuni/schubert-api/api/middleware"
)

// Score is model for score.
type Score struct {
	ID     bson.ObjectId `json:"id" bson:"_id,omitempty"`
	UserID bson.ObjectId `json:"userId" bson:"user_id,omitempty"`
	Sheets []struct {
		ScoreID bson.ObjectId `json:"scoreId" bson:"scoreId"`
		ID      string        `json:"id" bson:"id"`
		Index   int           `json:"index" bson:"index"`
	} `json:"sheets"`
	Staves []struct {
		SheetID string `json:"sheetId" bson:"sheetId"`
		ID      string `json:"id"`
		Index   int    `json:"index"`
	} `json:"staves"`
	Measures []struct {
		StaffID string `json:"staffId" bson:"staffId"`
		ID      string `json:"id"`
		Index   int    `json:"index"`
	} `json:"measures"`
	Voices []struct {
		MeasureID string `json:"measureId" bson:"measureId"`
		ID        string `json:"id"`
		Index     int    `json:"index"`
	} `json:"voices"`
	Notes []struct {
		VoiceID  string `json:"voiceId" bson:"voiceId"`
		ID       string `json:"id"`
		Index    int    `json:"index"`
		Value    string `json:"value"`
		Duration string `json:"duration"`
	} `json:"notes"`
	TimeSignatures []struct {
		MeasureID string `json:"measureId" bson:"measureId"`
		ID        string `json:"id"`
		Index     int    `json:"index"`
	} `json:"timeSignatures"`
	Clefs []struct {
		MeasureID string `json:"measureId" bson:"measureId"`
		ID        string `json:"id"`
		Index     int    `json:"index"`
	} `json:"clefs"`
}

// ListScores lists all scores for user.
func ListScores() http.Handler {
	return appHandler(func(w http.ResponseWriter, r *http.Request) *apiError {
		vars := mux.Vars(r)
		userID := vars["userID"]
		if !bson.IsObjectIdHex(userID) {
			return badInputError("Invalid id.")
		}

		var scores []Score
		session := middleware.MgoSessionFromContext(r.Context())
		c := session.DB("schubert").C("scores")
		query := map[string]bson.ObjectId{"user_id": bson.ObjectIdHex(userID)}
		if err := c.Find(query).All(&scores); err != nil {
			return &apiError{Code: 500, Message: "wtf", Error: err}
		}

		data := struct {
			Data []Score `json:"data"`
		}{Data: scores}
		js, err := json.Marshal(data)
		if err != nil {
			return internalServerError(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
		return nil
	})
}

// GetScore fetches score by id for user.
func GetScore() http.Handler {
	return appHandler(func(w http.ResponseWriter, r *http.Request) *apiError {
		vars := mux.Vars(r)
		id := vars["id"]
		if !bson.IsObjectIdHex(id) {
			return badInputError("Invalid id.")
		}

		userID := vars["userID"]
		if !bson.IsObjectIdHex(userID) {
			return badInputError("Invalid user id.")
		}

		session := middleware.MgoSessionFromContext(r.Context())

		var result Score
		query := map[string]bson.ObjectId{"_id": bson.ObjectIdHex(id), "user_id": bson.ObjectIdHex(userID)}
		if err := session.DB("schubert").C("scores").Find(query).One(&result); err != nil {
			if err.Error() == mgo.ErrNotFound.Error() {
				return notFoundError(err)
			}
			return internalServerError(err)
		}

		data := struct {
			Data Score `json:"data"`
		}{Data: result}
		js, err := json.Marshal(data)
		if err != nil {
			return internalServerError(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
		return nil
	})
}

// CreateScore creates a score.
func CreateScore() http.Handler {
	return appHandler(func(w http.ResponseWriter, r *http.Request) *apiError {
		vars := mux.Vars(r)
		userID := vars["userID"]
		if !bson.IsObjectIdHex(userID) {
			return badInputError("Invalid user id.")
		}

		var s Score
		err := json.NewDecoder(r.Body).Decode(&s)
		if err != nil {
			return internalServerError(err)
		}

		// TODO: Validate presence of userId.

		s.ID = bson.NewObjectId()
		for i := range s.Sheets {
			s.Sheets[i].ScoreID = s.ID
		}

		session := middleware.MgoSessionFromContext(r.Context())
		err = session.DB("schubert").C("scores").Insert(s)
		if err != nil {
			return internalServerError(err)
		}

		js, err := json.Marshal(s)
		if err != nil {
			return internalServerError(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(js)
		return nil
	})
}

// UpdateScore updates a score by id.
func UpdateScore() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})
}

// DeleteScore deletes a score by ID.
func DeleteScore() http.Handler {
	return appHandler(func(w http.ResponseWriter, r *http.Request) *apiError {
		vars := mux.Vars(r)
		userID := vars["userID"]
		if !bson.IsObjectIdHex(userID) {
			return badInputError("Invalid user id.")
		}

		id := vars["id"]
		if !bson.IsObjectIdHex(id) {
			return badInputError("Invalid id.")
		}

		session := middleware.MgoSessionFromContext(r.Context())
		err := session.DB("schubert").C("scores").RemoveId(bson.ObjectIdHex(id))
		if err != nil {
			return internalServerError(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		return nil
	})
}
