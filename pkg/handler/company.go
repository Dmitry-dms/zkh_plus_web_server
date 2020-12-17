package handler

import (
	"github.com/dmitry-dms/rest-gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type getAllCompaniesResponse struct {
	Data []models.Company `json:"data"`
}

func (h *Handler) getAllCompanies(c *gin.Context) {

	companyList, err := h.services.GetAllCompanies()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllCompaniesResponse{
		Data: companyList,
	})

}

type GetCompanyResponse struct {
	Data models.Company `json:"data"`
}

func (h *Handler) getCompanyById(c *gin.Context) {
	companyId, err1 := strconv.Atoi(c.DefaultQuery("company_id", "0"))
	if err1 != nil {
		newErrorResponse(c, http.StatusInternalServerError, "company_id is not type of int")
		return
	}
	if companyId == 0 || companyId < 0 {
		newErrorResponse(c, http.StatusInternalServerError, "company_id must be more than 0")
		return
	}
	company, err := h.services.GetCompanyById(companyId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, GetCompanyResponse{
		Data: company,
	})

}

func (h *Handler) createNotification(c *gin.Context) {
	var notification models.Notification
	companyId, err := getUserId(c)
	if err != nil {
		return
	}

	if err := c.BindJSON(&notification); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "input data are not of type notification")
		return
	}
	err = h.services.CompanyList.CreateNotification(companyId, notification)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, &successResponse{
			Message: "success",
		})
	}
}
