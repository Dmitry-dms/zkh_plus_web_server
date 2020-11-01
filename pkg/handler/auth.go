package handler

import (
	"github.com/dmitry-dms/rest-gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var input models.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error()) //400 - неверные данные в запросе
		return
	}
	//далее передаём данные в сервис,(слой ниже) где реализована бизнес логика регистрации

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error()) // 500 - внутренняя ошибка на сервере
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error()) //400 - неверные данные в запросе
		return
	}
	//далее передаём данные в сервис,(слой ниже) где реализована бизнес логика регистрации

	token, err := h.services.Authorization.GenerateToken(input.Email, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error()) // 500 - внутренняя ошибка на сервере
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
