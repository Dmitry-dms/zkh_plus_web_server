package handler

import (
	"github.com/dmitry-dms/rest-gin/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	test := router.Group("/")
	{
		test.GET("/t", h.testing)
		test.POST("/check-token", h.checkToken)
	}
	api := router.Group("/api/v1", h.userIdentity)
	{
		company := api.Group("/companies")
		{
			company.GET("/get-all", h.getAllCompanies)
			company.GET("/get", h.getCompanyById) // ?company_id=...
		}
		users := api.Group("/users")
		{
			users.POST("/update-company", h.updateUserCompany) // ?company_id=...
			users.POST("/create-address", h.createUserAddress)
			users.GET("/get-user-address", h.getAllUserAddress)
			users.POST("/insert-values", h.addVolumes)
			users.GET("/values/get-all", h.getAllUserValues)
			users.GET("/values/get", h.getUsersValuesByYearAndMonth) //?year=...&month=...
		}
	}
	return router
}
