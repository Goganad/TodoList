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

	listId, err := strconv.Atoi(mux.Vars(r)[idString])
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
		idString: id,
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

	listId, err := strconv.Atoi(mux.Vars(r)[idString])
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
	userId, err := getUserId(w, r.Context())
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(mux.Vars(r)[idString])
	if err != nil {
		return
	}

	item, err := h.services.TodoItem.GetById(userId, itemId)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, item)
}

func (h *Handler) updateItem(w http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(w, r.Context())
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(mux.Vars(r)[idString])
	if err != nil {
		return
	}

	var input entities.UpdateItemInput

	err = parseJsonToStruct(r, &input)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.services.TodoItem.Update(userId, itemId, input); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, statusResponse{
		Status: successResponse,
	})
}

func (h *Handler) deleteItem(w http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(w, r.Context())
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(mux.Vars(r)[idString])
	if err != nil {
		return
	}

	err = h.services.TodoItem.Delete(userId, itemId)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, statusResponse{
		Status: successResponse,
	})
}
