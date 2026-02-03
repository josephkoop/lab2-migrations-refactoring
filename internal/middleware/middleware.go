package middleware

import(
	"log"
	"net/http"
	"time"
)

var counter int = 0

//This middleware accepts the request and logs its method, path, and the current time once the request is received.
func LoggingMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		log.Printf("\nRequest method: %s \nRequest path: %s \nRequested at: %s\n", r.Method, r.URL.Path, time.Now().Local())
		next.ServeHTTP(w, r)
	})
}

//This middleware uses a simple global variable to keep a counter of requests received since server up time.
func CountingMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		counter += 1
		log.Printf("\nRequests since server up: %d\n\n", counter)
		next.ServeHTTP(w, r)
	})
}


