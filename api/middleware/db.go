package middleware

import (
	"context"
	"net/http"

	"github.com/globalsign/mgo"
)

type key int

const dbKey key = 1

// CopyMgoSession adds mongo session to context.
func CopyMgoSession(s *mgo.Session) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), dbKey, s.Copy())
			h.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// DBFromContext retrieves mongo session from context.
func DBFromContext(ctx context.Context) *mgo.Session {
	return ctx.Value(dbKey).(*mgo.Session)
}
