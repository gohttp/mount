package mount

import "net/http"
import "strings"

// Middleware constructor.
type Constructor func(http.Handler) http.Handler

// New mount middleware.
func New(prefix string, i interface{}) func(http.Handler) http.Handler {
	mw := middleware(i)
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			original := r.URL.Path

			if s := strings.TrimPrefix(r.URL.Path, prefix); len(s) < len(r.URL.Path) {
				r.URL.Path = s
				mw(h).ServeHTTP(w, r)
				r.URL.Path = original
			} else {
				h.ServeHTTP(w, r)
			}
		})
	}
}

// Coerce into middleware.
func middleware(h interface{}) func(http.Handler) http.Handler {
	switch h.(type) {
	case func(w http.ResponseWriter, r *http.Request):
		h := http.HandlerFunc(h.(func(w http.ResponseWriter, r *http.Request)))
		return func(_ http.Handler) http.Handler {
			return http.HandlerFunc(h)
		}
	case http.Handler:
		h := h.(http.Handler)
		return func(_ http.Handler) http.Handler {
			return h
		}
	case func(h http.Handler) http.Handler:
		return h.(func(h http.Handler) http.Handler)
	default:
		panic("invalid middleware")
	}
}
