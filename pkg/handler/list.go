package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createList(c *gin.Context) {
	id, _ := c.Get("userId")
	c.JSON(http.StatusOK, map[string]interface{}{
		"userId": id,
	})
}

func (h *Handler) getAllList(c *gin.Context) {

}

func (h *Handler) getListById(c *gin.Context) {

}

func (h *Handler) updateList(c *gin.Context) {

}

func (h *Handler) deleteList(c *gin.Context) {

}
