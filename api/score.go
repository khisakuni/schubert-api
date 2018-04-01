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
	Sheets []struct {
		ScoreID bson.ObjectId `json:"scoreId" bson:"scoreId"`
		ID      string        `json:"id" bson:"id"`
		Index   int           `json:"index" bson:"index"`
	}
	Staves []struct {
		SheetID string `json:"sheetId" bson:"sheetId"`
		ID      string `json:"id"`
		Index   int    `json:"index"`
	}
	Measures []struct {
		StaffID string `json:"staffId" bson:"staffId"`
		ID      string `json:"id"`
		Index   int    `json:"index"`
	}
	Voices []struct {
		MeasureID string `json:"measureId" bson:"measureId"`
		ID        string `json:"id"`
		Index     int    `json:"index"`
	}
	Notes []struct {
		VoiceID  string `json:"voiceId" bson:"voiceId"`
		ID       string `json:"id"`
		Index    int    `json:"index"`
		Value    string `json:"value"`
		Duration string `json:"duration"`
	}
	TimeSignatures []struct {
		MeasureID string `json:"measureId" bson:"measureId"`
		ID        string `json:"id"`
		Index     int    `json:"index"`
	}
	Clefs []struct {
		MeasureID string `json:"measureId" bson:"measureId"`
		ID        string `json:"id"`
		Index     int    `json:"index"`
	}
}

// ListScores lists all scores.
func ListScores() http.Handler {
	return appHandler(func(w http.ResponseWriter, r *http.Request) *apiError {
		var scores []Score
		session := middleware.MgoSessionFromContext(r.Context())
		c := session.DB("schubert").C("scores")
		if err := c.Find(nil).All(&scores); err != nil {
			return &apiError{Code: 500, Message: "wtf", Error: err}
		}

		js, err := json.Marshal(scores)
		if err != nil {
			return internalServerError(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
		return nil
	})
}

// GetScore fetches score by id.
func GetScore() http.Handler {
	return appHandler(func(w http.ResponseWriter, r *http.Request) *apiError {
		vars := mux.Vars(r)
		id := vars["id"]
		if !bson.IsObjectIdHex(id) {
			return badInputError("Invalid id.")
		}

		session := middleware.MgoSessionFromContext(r.Context())

		var result Score
		if err := session.DB("schubert").C("scores").FindId(bson.ObjectIdHex(id)).One(&result); err != nil {
			if err.Error() == mgo.ErrNotFound.Error() {
				return notFoundError(err)
			}
			return internalServerError(err)
		}

		js, err := json.Marshal(result)
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
		var s Score
		err := json.NewDecoder(r.Body).Decode(&s)
		if err != nil {
			return internalServerError(err)
		}

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
		w.WriteHeader(201)
		w.Write(js)
		return nil
	})
}

// UpdateScore updates a score by id.
func UpdateScore() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})
}
