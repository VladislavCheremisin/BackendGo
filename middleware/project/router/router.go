package router

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() *mux.Router {
	// mux.Router
	router := mux.NewRouter()
	// Items from list Items with filters
	router.HandleFunc("/items", ListItemsHandler).Methods(http.MethodGet)
	// new Item
	router.HandleFunc("/items", CreateItemHandler).Methods(http.MethodPost)
	//  Get Item
	router.HandleFunc("/items/{id}", GetItemHandler).Methods(http.MethodGet)
	// Update Item
	router.HandleFunc("/items/{id}", UpdateItemHandler).Methods(http.MethodPut)
	//  Delete Item  ID
	router.HandleFunc("/items/{id}", DeleteItemHandler).Methods(http.MethodDelete)
	// Upload ItemI mage
	router.HandleFunc("/items/upload_image", UploadItemImageHandler).Methods(http.MethodPost)
	// Login
	router.HandleFunc("/user/login", LoginHandler).Methods(http.MethodPost)
	// Logout
	router.HandleFunc("/user/logout", LogoutHandler).Methods(http.MethodPost)
	return router
}
