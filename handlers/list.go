package handlers

import (
	"net/http"
)

func (h *Handler) createList(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(userCtx)

}

func (h *Handler) getAllLists(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) getListById(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) updateList(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) deleteList(w http.ResponseWriter, r *http.Request) {

}
