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

	r.HandleFunc("/auth/signin", h.SignIn).Methods("POST")
	r.HandleFunc("/auth/signup", h.SignUp).Methods("POST")
	r.HandleFunc("/api/lists/", h.GetAllLists).Methods("GET")
	r.HandleFunc("/api/lists/", h.CreateList).Methods("POST")
	r.HandleFunc("/api/lists/{id}", h.GetListById).Methods("GET")
	r.HandleFunc("/api/lists/{id}", h.UpdateList).Methods("PUT")
	r.HandleFunc("/api/lists/{id}", h.DeleteList).Methods("DELETE")
	r.HandleFunc("/api/lists/{id}/items", h.GetAllItems).Methods("GET")
	r.HandleFunc("/api/lists/{id}/items", h.CreateItem).Methods("POST")
	r.HandleFunc("/api/lists/{id}/items/{item_id}", h.GetItemById).Methods("GET")
	r.HandleFunc("/api/lists/{id}/items/{item_id}", h.UpdateItem).Methods("PUT")
	r.HandleFunc("/api/lists/{id}/items/{item_id}", h.DeleteItem).Methods("DELETE")

	return r
}