package handlers

import (
	"github.com/Goganad/TodoList-REST-API/entities"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) createItem(w http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(w, r.Context())
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return
	}

	var input entities.TodoItem

	err = parseJsonToStruct(r, &input)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusBadRequest, "Invalid input")
		return
	}

	id, err := h.services.TodoItem.Create(userId, listId, input)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type itemsResponse struct {
	Data []entities.TodoItem `json:"data"`
}

func (h *Handler) getAllItems(w http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(w, r.Context())
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return
	}

	items, err := h.services.TodoItem.GetAll(userId, listId)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, itemsResponse{
		Data: items,
	})
}

func (h *Handler) getItemById(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) updateItem(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) deleteItem(w http.ResponseWriter, r *http.Request) {

}
