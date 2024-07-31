package handler

import (
	"deliver/internal/service"

	"deliver/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.HandleMethodNotAllowed = true
	router.Use(corsMiddleware())

	//swagger settings
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler), func(c *gin.Context) {
		docs.SwaggerInfo.Host = c.Request.Host
		if c.Request.TLS != nil {
			docs.SwaggerInfo.Schemes = []string{"https"}
		}
	})

	// AUTH
	router.POST("api/v1/auth/login", h.login)
	router.POST("api/v1/auth/refresh", h.refresh)

	api := router.Group("/api/v1", h.userIdentity)
	{
		minio := api.Group("/minio")
		{
			minio.POST("/upload-image", h.uploadImage)
		}

		user := api.Group("/users")
		{
			user.POST("", h.createUser)
		}

		role := api.Group("/roles")
		{
			role.GET("", h.getListRole)
		}

		category := api.Group("/categories")
		{
			category.POST("", h.createCategory)
			category.GET("", h.getListCategory)
			category.GET("/:id", h.getCategoryById)
			category.PUT("/:id", h.updateCategory)
			category.DELETE("/:id", h.deleteCategory)
		}

		product := api.Group("/products")
		{
			product.POST("", h.createProduct)
			product.GET("", h.getListProduct)
			product.GET("/:id", h.getProductById)
			product.PUT("/:id", h.updateProduct)
			product.DELETE("/:id", h.deleteProduct)
		}

		attribute := api.Group("/attributes")
		{
			attribute.POST("", h.createAttribute)
			attribute.GET("", h.getListAttribute)
			attribute.GET("/:id", h.getAttributeById)
			attribute.PUT("/:id", h.updateAttribute)
			attribute.DELETE("/:id", h.deleteAttribute)
		}

		option := attribute.Group("/:id/options")
		{
			option.POST("", h.createOption)
			option.GET("", h.getListOption)
			option.GET("/:option-id", h.getOptionById)
			option.PUT("/:option-id", h.updateOption)
			option.DELETE("/:option-id", h.deleteOption)
		}

	}

	return router
}
