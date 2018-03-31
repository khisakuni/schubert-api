package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/khisakuni/schubert-api/api/middleware"
	"gopkg.in/mgo.v2/bson"
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
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})
}

// GetScore fetches score by id.
func GetScore() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session := middleware.MgoSessionFromContext(r.Context())
		fmt.Printf("Session from context >> %v", session)
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

		// Create score
		scoreID := bson.NewObjectId()
		s.ID = scoreID

		// Create sheets
		for i := range s.Sheets {
			s.Sheets[i].ScoreID = scoreID
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
