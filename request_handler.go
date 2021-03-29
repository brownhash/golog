package golog

import (
	"fmt"
	"log"
	"net/http"
)

// LogRequest - log http requests
func LogRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		formatter := fmt.Sprintf("%s %s", fmt.Sprintf(SuccessColor, r.Method), fmt.Sprintf(DebugColor, r.URL))
		log.Printf(formatter)

		handler.ServeHTTP(w, r)
	})
}

/*
Usage-

http.ListenAndServe("127.0.0.1:8080", golog.LogRequest(http.DefaultServeMux))
*/