package handlers

import (
	"github.com/Goganad/TodoList-REST-API/entities"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) createList(w http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(w, r.Context())
	if err != nil {
		return
	}

	var input entities.TodoList

	err = parseJsonToStruct(r, &input)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusBadRequest, "Invalid input")
		return
	}

	id, err := h.services.TodoList.Create(userId, input)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type listsResponse struct {
	Data []entities.TodoList `json:"data"`
}

func (h *Handler) getAllLists(w http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(w, r.Context())
	if err != nil {
		return
	}

	lists, err := h.services.TodoList.GetAll(userId)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, listsResponse{
		Data: lists,
	})
}

func (h *Handler) getListById(w http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(w, r.Context())
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return
	}

	list, err := h.services.TodoList.GetById(userId, listId)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, list)
}

func (h *Handler) updateList(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) deleteList(w http.ResponseWriter, r *http.Request) {

}
