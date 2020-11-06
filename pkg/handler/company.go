package handler

import (
	"fmt"
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
	if companyList == nil {
		fmt.Printf("list = %s", companyList)
		return
	}

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
