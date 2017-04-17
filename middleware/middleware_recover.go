package middleware

import (
	"runtime"
	"net/http"
	"github.com/CardInfoLink/log"
)

func RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Debug("[middleware] enter panic recover")
		defer func() {
			if e := recover(); e != nil {
				trace := make([]byte, 1024)
				count := runtime.Stack(trace, false)
				log.Errorf("%v stacktrace: %s\n", e, trace[:count])
			}
		}()
		next.ServeHTTP(w, r)
		log.Debug("[middleware] exit panic recover")
	});
}