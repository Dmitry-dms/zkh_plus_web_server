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
	}
	api := router.Group("/api/v1", h.userIdentity)
	{
		company := api.Group("/users")
		{
			company.GET("/", h.getAllCompanies)
			//company.GET("/:id", h.getUserById)
		}
	}
	return router
}
