package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) testing(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"test": "Тест пройден",
	})
}