package handlers

import (
	"context"
	"errors"
	"net/http"
	"strings"
)

type contextKey string

const (
	validHeadersCount              = 2
	authorizationHeader            = "Authorization"
	userCtx             contextKey = "userId"
)

func (h *Handler) isAuthenticated(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get(authorizationHeader)
		if header == "" {
			respondWithError(w, http.StatusUnauthorized, "empty auth header")
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != validHeadersCount {
			respondWithError(w, http.StatusUnauthorized, "invalid auth header")
			return
		}

		userId, err := h.services.Authorization.ParseToken(headerParts[1])
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, err.Error())
			return
		}

		ctx := context.WithValue(r.Context(), userCtx, userId)
		next(w, r.WithContext(ctx))
	})
}

func getUserId(w http.ResponseWriter, ctx context.Context) (int, error) {
	userId := ctx.Value(userCtx)

	id, ok := userId.(int)
	if !ok {
		respondWithError(w, http.StatusInternalServerError, "user id not found")
		return 0, errors.New("user id not found")
	}

	return id, nil
}
