package route

import (
	"github.com/cli-server/api"
	"github.com/cli-server/db"
	"github.com/cli-server/handler"
	myMw "github.com/cli-server/middleware"
	"github.com/labstack/echo"
	echoMw "github.com/labstack/echo/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	// Debug mode
	e.Debug = true

	// Bundle middleware
	//-------------------
	e.Use(echoMw.LoggerWithConfig(echoMw.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, latency=${latency_human}, time=${time_rfc3339}\n",
	}))
	e.Use(echoMw.Gzip())
	e.Use(echoMw.CORSWithConfig(echoMw.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding},
	}))

	// Custom middleware
	//-------------------
	e.HTTPErrorHandler = handler.JSONHTTPErrorHandler
	// Stats
	s := myMw.NewStats()
	e.Use(s.Process)
	// TransactionHandler
	e.Use(myMw.TransactionHandler(db.Init()))

	// Server header
	e.Use(myMw.ServerHeader)

	// Routes
	//-------------------
	e.GET("/stats", s.Handle) // Endpoint to get stats

	// Users Collection
	users := e.Group("/api/users")
	{
		// Creat/Update an User
		users.POST("", api.PostUser())
		// Get an User
		users.GET("", api.GetUsers())
		users.GET("/:open_id", api.GetUser())
	}

	//Merchants Collection
	merchants := e.Group("/api/merchants")
	{
    //post a merchants
    merchants.POST("", api.PostMerchant())
		//Get all Merchants
		merchants.GET("", api.GetAllMerchant())
		//get merchant with id
		merchants.GET("/:merchant_id", api.GetMerchant())
		//TODO: update merchant with id
	}
	foods := e.Group("/api/foods")
	{
		//Get all Foods by MerchantID
		foods.GET("/:merchan_id", api.GetFoodsByMerchantId())
		// Create a New Food,with a merchant-id
		foods.POST("", api.PostFood())
		// Get a Food by FoodID
		foods.GET("/:food_id", api.GetFoodByFoodId())
		// Update a Food by FoodID
		foods.POST("", api.PostFood())
		//Delete a Food by FoodID
		foods.DELETE("/:food_id", api.DeleteFood())
	}
	orders := e.Group("/api/orders")
	{
		//Orders / Get all Orders by OpenID
		orders.GET("/:open_id", api.GetOrdersByOpenId())
		//Orders / Get all Orders by MerchantID
		orders.GET("/:merchant_id", api.GetOrdersByMerchant())
		//Orders / Get an Order by OrderID
		orders.GET("/:order_id", api.GetOrdersByOrderId())
		//create/ Update an Order by OrderID
		orders.POST("", api.PostOrder())
		//Orders / Delete an Order by OrderID
		orders.DELETE("/:order_id", api.DeleteOrderByOrderId())

	}

	return e
}
