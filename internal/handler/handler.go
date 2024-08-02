package handler

import (
	"deliver/internal/service"
	"deliver/internal/ws"

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
	router.POST("api/v1/auth/signup", h.signUp)
	router.POST("api/v1/auth/refresh", h.refresh)

	// WebSocket route
	router.GET("/ws", func(c *gin.Context) {
		ws.HandleWebSocket(c.Writer, c.Request)
	})

	api := router.Group("/api/v1", h.userIdentity)
	{
		minio := api.Group("/minio")
		{
			minio.POST("/upload-image", h.uploadImage)
		}

		user := api.Group("/users")
		{
			user.GET("/me", h.getUserMe)
		}

		category := api.Group("/categories")
		{
			category.GET("", h.getListCategory)
			category.GET("/:id", h.getCategoryById)
		}

		product := api.Group("/products")
		{
			product.GET("", h.getListProduct)
			product.GET("/:id", h.getProductById)
		}

		attribute := api.Group("/attributes")
		{
			attribute.GET("", h.getListAttribute)
			attribute.GET("/:id", h.getAttributeById)
		}

		option := attribute.Group("/:id/options")
		{
			option.GET("", h.getListOption)
			option.GET("/:option-id", h.getOptionById)
		}

		order := api.Group("/orders")
		{
			order.GET("", h.getListOrder)
			order.GET("/:id", h.getOrderById)
		}

		notification := api.Group("/notifications")
		{
			notification.POST("/customer", h.sendNotificationToCustomer)
			notification.POST("/admin", h.sendNotificationToAdmin)
			notification.POST("/courier", h.sendNotificationToCourier)
		}
	}

	admin := api.Group("", isAdminMiddleware())
	{
		user := admin.Group("/users")
		{
			user.POST("", h.createUser)
		}

		role := admin.Group("/roles")
		{
			role.GET("", h.getListRole)
		}

		category := admin.Group("/categories")
		{
			category.POST("", h.createCategory)
			category.PUT("/:id", h.updateCategory)
			category.DELETE("/:id", h.deleteCategory)
		}

		product := admin.Group("/products")
		{
			product.POST("", h.createProduct)
			product.PUT("/:id", h.updateProduct)
			product.DELETE("/:id", h.deleteProduct)
			product.POST("/:id/add/:attribute-id", h.addAttributeToProduct)
			product.DELETE("/:id/remove/:attribute-id", h.removeAttributeFromProduct)
		}

		attribute := admin.Group("/attributes")
		{
			attribute.POST("", h.createAttribute)
			attribute.PUT("/:id", h.updateAttribute)
			attribute.DELETE("/:id", h.deleteAttribute)
		}

		option := attribute.Group("/:id/options")
		{
			option.POST("", h.createOption)
			option.PUT("/:option-id", h.updateOption)
			option.DELETE("/:option-id", h.deleteOption)
		}
	}

	courier := api.Group("", isCourierMiddleware())
	{
		order := courier.Group("/orders")
		{
			order.POST("/:id/receive-courier", h.receiveOrderCourier)
			order.POST("/:id/finish-courier", h.finishOrderCourier)
			order.POST("/:id/payment-collect", h.paymentCollectOrderCourier)
		}
	}

	customer := api.Group("", isCustomerMiddleware())
	{
		order := customer.Group("/orders")
		{
			order.POST("", h.createOrder)
			order.GET("/history", h.getOrderHistory)
		}
	}

	ws.StartWebSocketHub() // Start the WebSocket hub

	return router
}
