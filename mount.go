package mount

import "net/http"
import "strings"

// Middleware constructor.
type Constructor func(http.Handler) http.Handler

// New mount middleware.
func New(prefix string, mw Constructor) func(http.Handler) http.Handler {
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
