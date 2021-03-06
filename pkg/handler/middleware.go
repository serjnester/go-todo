package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
)

func (h *Handler) userIdentity(c *gin.Context) {
	//отримуємо хедер авторизації
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}
	//розбиваємо хедер на частини
	headersParts := strings.Split(header, " ")
	if len(headersParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	userId, err := h.services.Authorization.ParseToken(headersParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
	}

	c.Set("userId", userId)
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get("userId")
	if !ok {
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("user id is of invalid type")
	}

	return idInt, nil
}
