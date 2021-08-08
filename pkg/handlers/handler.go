package handlers

import (
	"github.com/Goganad/TodoList-REST-API/pkg/service"
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
	r.Handle("/api/lists", h.isAuthenticated(h.getAllLists)).Methods("GET")
	r.Handle("/api/lists", h.isAuthenticated(h.createList)).Methods("POST")
	r.Handle("/api/lists/{id}", h.isAuthenticated(h.getListById)).Methods("GET")
	r.Handle("/api/lists/{id}", h.isAuthenticated(h.updateList)).Methods("PUT")
	r.Handle("/api/lists/{id}", h.isAuthenticated(h.deleteList)).Methods("DELETE")
	r.Handle("/api/lists/{id}/items", h.isAuthenticated(h.getAllItems)).Methods("GET")
	r.Handle("/api/lists/{id}/items", h.isAuthenticated(h.createItem)).Methods("POST")
	r.Handle("/api/items/{id}", h.isAuthenticated(h.getItemById)).Methods("GET")
	r.Handle("/api/items/{id}", h.isAuthenticated(h.updateItem)).Methods("PUT")
	r.Handle("/api/items/{id}", h.isAuthenticated(h.deleteItem)).Methods("DELETE")

	return r
}
