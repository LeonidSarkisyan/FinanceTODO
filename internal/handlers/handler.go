package handlers

import (
	"FinanceTODO/internal/services"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *services.Service
}

func NewHandler(services *services.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/login", h.Login)
		auth.POST("/register", h.Register)
	}

	api := router.Group("/api", h.CurrentUser)
	{
		users := api.Group("/users", h.CurrentUser)
		{
			users.GET("/me", h.GetUserById)
		}

		subCategories := api.Group("/subcategories")
		{
			subCategories.POST("/", h.CreateSubCategory)
			subCategories.GET("/", h.GetSubCategories)
			subCategories.GET("/:id", h.GetSubCategoryById)
			subCategories.PATCH("/:id", h.UpdateSubCategory)
			subCategories.DELETE("/:id", h.DeleteSubCategory)
		}

		balances := api.Group("/balances")
		{
			balances.POST("/", h.CreateBalance)
			balances.GET("/", h.GetBalances)
			balances.GET("/:id", h.GetBalanceById)
			balances.PATCH("/:id", h.UpdateBalance)
			balances.DELETE("/:id", h.DeleteBalance)

			balances = balances.Group(":id")

			transactions := balances.Group("/transactions")
			{
				transactions.POST("/", h.CreateTransaction)
				transactions.GET("/", h.GetTransactions)
				transactions.GET("/:transaction_id", h.GetTransactionById)
				transactions.PATCH("/:transaction_id", h.UpdateTransaction)
				transactions.DELETE("/:transaction_id", h.DeleteTransaction)
			}
		}
	}

	admin := router.Group("/admin")
	{
		users := admin.Group("/users")
		{
			users.POST("/", h.Register)
			//users.GET("/")
			//users.GET("/:user_id")
			//users.PATCH("/:user_id")
			//users.DELETE("/:user_id")
		}
	}
	//
	//api := router.Group("/api")
	//{
	//	api.POST("/login")
	//	api.POST("/register")
	//}

	return router
}
