package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("request: %s %s - %v\n",
			r.Method,
			r.URL.EscapedPath(),
			time.Since(start),
		)
	})
}

func RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if e := recover(); e != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintln(w, "we've got panic here!")
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func main() {
	// объявляем mux.Router
	router := chi.NewRouter()

	router.Use(LoggingMiddleware)
	router.Use(RecoverMiddleware)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if e := recover(); e != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintln(w, "we've got panic here!")
			}
		}()
		panic("panic!")
		fmt.Fprintln(w, "GET HANDLER")
	})

	router.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		fmt.Fprintln(w, "GET BY ID HANDLER. RESOURCE ID IS", id)
		return
	})

	router.Get("/{id}/name/{name}", func(w http.ResponseWriter, r *http.Request) {
		id, name := chi.URLParam(r, "id"), chi.URLParam(r, "name")
		fmt.Fprintf(w, "GET BY ID HANDLER WITH NAME. RESOURCE ID IS %s AND NAME IS %s\n", id, name)
		return
	})

	log.Fatal(http.ListenAndServe(":8020", router))
}
