package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header") //401 - не авторизирован
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	//парсим токен
	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
	}

	c.Set(userCtx, userId)
}

//func (h *Handler) companyIdentity(c *gin.Context) {
//	header := c.GetHeader(authorizationHeader)
//	if header == "" {
//		newErrorResponse(c, http.StatusUnauthorized, "empty auth header") //401 - не авторизирован
//		return
//	}
//	headerParts := strings.Split(header, " ")
//	if len(headerParts) != 2 {
//		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
//		return
//	}
//
//	//парсим токен
//	companyId, err := h.services.Authorization.ParseToken(headerParts[1])
//	if err != nil {
//		newErrorResponse(c, http.StatusUnauthorized, err.Error())
//	}
//
//	c.Set(userCtx, companyId)
//}
func (h *Handler) checkToken(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header") //401 - не авторизирован
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}
	//парсим токен
	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
	} else {

		c.JSON(http.StatusOK, map[string]interface{}{
			"user_id": userId,
		})
	}
}

func getUserId(c *gin.Context) (primitive.ObjectID, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user_id not found")
		return primitive.NilObjectID, errors.New("user_id not found")
	}
	idInt, ok := id.(primitive.ObjectID)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "invalid type of user_id")
		return primitive.NilObjectID, errors.New("invalid type of user_id")
	}
	return idInt, nil
}
