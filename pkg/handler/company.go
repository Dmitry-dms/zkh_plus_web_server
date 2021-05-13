package handler

import (
	"github.com/Dmitry-dms/zkh-plus/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Company struct {
	DirectorFullName string `json:"director_fio"`
	Phone            string `json:"company_phone"`
	City             string `json:"company_city"`
	Street           string `json:"company_street"`
	HomeNumber       string `json:"company_home_number"`
	Flat             string `json:"company_flat"`
	Name             string `json:"company_name"`
}
type getAllCompaniesResponse struct {
	Data []Company `json:"data"`
}

func (h *Handler) getAllCompanies(c *gin.Context) {

	companyList, err := h.services.GetAllCompanies()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var list []Company
	for _,c := range companyList {
		var s Company
		s.City=c.City
		s.DirectorFullName=c.DirectorName
		s.Flat=c.Flat
		s.HomeNumber=c.HomeNumber
		s.Phone=c.Phone
		s.Street=c.Street
		s.Name=c.Name
		list = append(list, s)
	}
	c.JSON(http.StatusOK, getAllCompaniesResponse{
		Data: list,
	})

}

type GetCompanyResponse struct {
	Data models.Company `json:"data"`
}

func (h *Handler) getCompanyById(c *gin.Context) {
	// companyId, err1 := strconv.Atoi(c.DefaultQuery("company_id", "0"))
	// if err1 != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError, "company_id is not type of int")
	// 	return
	// }
	// if companyId == 0 || companyId < 0 {
	// 	newErrorResponse(c, http.StatusInternalServerError, "company_id must be more than 0")
	// 	return
	// }
	// companyId := c.DefaultQuery("company_id", "0")
	// company, err := h.services.GetCompanyById(companyId)
	// if err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }
	// c.JSON(http.StatusOK, GetCompanyResponse{
	// 	Data: company,
	// })

}

func (h *Handler) createNotification(c *gin.Context) {
	// var notification models.Notification
	// companyId, err := getUserId(c)
	// if err != nil {
	// 	return
	// }

	// if err := c.BindJSON(&notification); err != nil {
	// 	newErrorResponse(c, http.StatusBadRequest, "input data are not of type notification")
	// 	return
	// }
	// err = h.services.CompanyList.CreateNotification(companyId, notification)
	// if err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// } else {
	// 	c.JSON(http.StatusOK, &successResponse{
	// 		Message: "success",
	// 	})
	// }
}
