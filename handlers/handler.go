package handlers

import (
	"github.com/Goganad/TodoList-REST-API/service"
	"github.com/gorilla/mux"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/auth/signin", h.signIn).Methods("POST")
	r.HandleFunc("/auth/signup", h.signUp).Methods("POST")
	r.HandleFunc("/api/lists/", h.getAllLists).Methods("GET")
	r.HandleFunc("/api/lists/", h.createList).Methods("POST")
	r.HandleFunc("/api/lists/{id}", h.getListById).Methods("GET")
	r.HandleFunc("/api/lists/{id}", h.updateList).Methods("PUT")
	r.HandleFunc("/api/lists/{id}", h.deleteList).Methods("DELETE")
	r.HandleFunc("/api/lists/{id}/items", h.getAllItems).Methods("GET")
	r.HandleFunc("/api/lists/{id}/items", h.createItem).Methods("POST")
	r.HandleFunc("/api/lists/{id}/items/{item_id}", h.getItemById).Methods("GET")
	r.HandleFunc("/api/lists/{id}/items/{item_id}", h.updateItem).Methods("PUT")
	r.HandleFunc("/api/lists/{id}/items/{item_id}", h.deleteItem).Methods("DELETE")

	return r
}
