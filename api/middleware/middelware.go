package middleware

import "net/http"

// Adapter is a wrapper for http.Handler funcs
type Adapter func(http.Handler) http.Handler

// Adapt applies all passed in Adapters.
func Adapt(h http.Handler, adapters ...Adapter) http.Handler {
	for _, adapter := range adapters {
		h = adapter(h)
	}
	return h
}
