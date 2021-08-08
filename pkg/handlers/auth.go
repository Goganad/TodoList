package handlers

import (
	"github.com/Goganad/TodoList-REST-API/entities"
	"log"
	"net/http"
)

type signInInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	var input signInInput

	err := parseJsonToStruct(r, &input)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusBadRequest, "Invalid input")
		return
	}

	tokenString, err := h.services.GenerateToken(input.Username, input.Password)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]interface{}{
		"token": tokenString,
	})
}

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	var input entities.User

	err := parseJsonToStruct(r, &input)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusBadRequest, "Invalid input")
		return
	}

	id, err := h.services.CreateUser(input)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
