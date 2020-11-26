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
		auth.POST("/sign-in", h.signIn)
	}
	test := router.Group("/")
	{
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
			users.GET("/notifications", h.getNotifications)          //?company_id=...
		}
	}
	companyOwner := router.Group("/c-own")
	{
		companyOwnerAuth := companyOwner.Group("/auth")
		{
			companyOwnerAuth.POST("/sign-up", h.companySignUp)
			companyOwnerAuth.POST("/sign-in", h.companySignIn)
		}
		companyOwnerRequests := companyOwner.Group("/req", h.userIdentity)
		{
			companyOwnerRequests.POST("/sign-up", h.signUp) //create user
			companyOwnerRequests.POST("/create-notification", h.createNotification)
		}

	}
	return router
}
