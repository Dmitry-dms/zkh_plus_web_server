package handler

import (
	"github.com/dmitry-dms/rest-gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) updateUserCompany(c *gin.Context) {

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	companyId, err1 := strconv.Atoi(c.DefaultQuery("company_id", "0"))
	if err1 != nil {
		newErrorResponse(c, http.StatusInternalServerError, "company_id is not type of int")
		return
	}
	if companyId == 0 {
		newErrorResponse(c, http.StatusInternalServerError, "user_id=0")
		return
	}
	err = h.services.UpdateUserCompany(userId, companyId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"update_company": "success",
		})
	}
}

func (h *Handler) createUserAddress(c *gin.Context) {
	var address models.UserAddress
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	if err := c.BindJSON(&address); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error()) //400 - неверные данные в запросе
		return
	}
	id, err := h.services.UserRequest.CreateUserAddress(userId, address)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error()) // 500 - внутренняя ошибка на сервере
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"address_id": id,
	})
}

type allUserAddressResponse struct {
	Data []models.UserAddress `json:"data"`
}

func (h *Handler) getAllUserAddress(c *gin.Context) {

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	addressList, err := h.services.GetAllUserAddress(userId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, allUserAddressResponse{
		Data: addressList,
	})
}
