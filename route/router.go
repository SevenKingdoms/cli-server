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
		// TODO: Use True Api
		// Creat/Update an User
		users.POST("", api.PostUser())
		// Get an User
		users.GET("", api.GetUsers())
		users.GET("/:open_id", api.GetUser())
	}

	//Merchants Collection
	merchants := e.Group("/api/merchants")
	{
		// TODO: Use True Api

		//Get all Merchants
		merchants.GET("", api.GetAllMerchant())
		//post a merchants
		merchants.POST("", api.PostMerchant())
		//get merchant with id
		merchants.GET("/:id", api.GetMerchant())
		//TODO: update merchant with id
	}

	return e
}
