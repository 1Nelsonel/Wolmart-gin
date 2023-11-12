package main

import (
	"github.com/gin-gonic/gin"

	database "github.com/1Nelsonel/Wolmart/Database"
	route "github.com/1Nelsonel/Wolmart/Route"
)

// Middleware to initialize db connections
func init() {
	database.ConnectDB()
}

func main() {
	// Database connect
	sqlDb, err := database.DBConn.DB()

	if err != nil {
		panic("Error in connection")
	}

	defer sqlDb.Close()
	app := gin.Default()

	// Templates
	app.LoadHTMLGlob("templates/*/*")

	// Serve static files (CSS, JavaScript, images) from the "public" directory
	app.Static("/assets", "./assets")

	// Initialize routes by calling SetupRoutes
	route.SetupRoutes(app)

	app.Run()

}
