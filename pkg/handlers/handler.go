package handlers

import (
	"github.com/Goganad/TodoList-REST-API/pkg/service"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/auth/signin", h.signIn).Methods(http.MethodPost)
	r.HandleFunc("/auth/signup", h.signUp).Methods(http.MethodPost)
	r.Handle("/api/lists", h.isAuthenticated(h.getAllLists)).Methods(http.MethodGet)
	r.Handle("/api/lists", h.isAuthenticated(h.createList)).Methods(http.MethodPost)
	r.Handle("/api/lists/{id}", h.isAuthenticated(h.getListById)).Methods(http.MethodGet)
	r.Handle("/api/lists/{id}", h.isAuthenticated(h.updateList)).Methods(http.MethodPut)
	r.Handle("/api/lists/{id}", h.isAuthenticated(h.deleteList)).Methods(http.MethodDelete)
	r.Handle("/api/lists/{id}/items", h.isAuthenticated(h.getAllItems)).Methods(http.MethodGet)
	r.Handle("/api/lists/{id}/items", h.isAuthenticated(h.createItem)).Methods(http.MethodGet)
	r.Handle("/api/items/{id}", h.isAuthenticated(h.getItemById)).Methods(http.MethodGet)
	r.Handle("/api/items/{id}", h.isAuthenticated(h.updateItem)).Methods(http.MethodPut)
	r.Handle("/api/items/{id}", h.isAuthenticated(h.deleteItem)).Methods(http.MethodDelete)

	return r
}
