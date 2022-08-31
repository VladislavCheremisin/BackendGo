package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	// объявляем mux.Router
	router := mux.NewRouter()

	router.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fmt.Fprintln(w, "GET BY ID HANDLER. RESOURCE ID IS", vars["id"])
		return
	}).Methods(http.MethodGet)

	router.HandleFunc("/{id}/name/{name}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fmt.Fprintf(w, "GET BY ID HANDLER WITH NAME. RESOURCE ID IS %s AND NAME IS %s\n",
			vars["id"], vars["name"])
		return
	}).Methods(http.MethodGet)

	// запускаем сервер, передав в качестве маршрутизатора объект mux.Router
	log.Fatal(http.ListenAndServe(":8010", router))
}
