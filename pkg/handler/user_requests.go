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
		c.JSON(http.StatusOK, &successResponse{
			Message: "success",
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
func (h *Handler) addVolumes(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	var volume models.DataVolume
	if err := c.BindJSON(&volume); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error()) //400 - неверные данные в запросе
		return
	}
	err = h.services.InputVolumes(userId, volume)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error()) //400 - неверные данные в запросе
		return
	}
	c.JSON(http.StatusOK, &successResponse{
		Message: "success",
	})
}

type userValuesResponse struct {
	Data []models.DataVolume `json:"data"`
}

func (h *Handler) getUsersValuesByYearAndMonth(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	year, err := strconv.Atoi(c.DefaultQuery("year", "0"))
	month, err1 := strconv.Atoi(c.DefaultQuery("month", "0"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "year is not type of int")
		return
	}
	if err1 != nil {
		newErrorResponse(c, http.StatusInternalServerError, "year is not type of int")
		return
	}
	if year == 0 || year < 0 {
		newErrorResponse(c, http.StatusInternalServerError, "year should be more than 0")
		return
	}
	if month == 0 || month < 0 {
		newErrorResponse(c, http.StatusInternalServerError, "month should be more than 0")
		return
	}

	list, err := h.services.GetUsersValuesByYearAndMonth(userId, year, month)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, userValuesResponse{
		Data: list,
	})
}
func (h *Handler) getAllUserValues(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	list, err := h.services.GetAllUserValues(userId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, userValuesResponse{
		Data: list,
	})
}
