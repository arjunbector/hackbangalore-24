package main

import (
	"context"

	"github.com/SohamGhugare/hackbangalore-24/database"
	"github.com/SohamGhugare/hackbangalore-24/handlers"
	"github.com/SohamGhugare/hackbangalore-24/initializers"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func init() {
	// loading dotenv file
	initializers.LoadEnv()

	// connecting to database
	database.ConnectDB()

}

func setupRoutes(r *gin.Engine) {
	// INVESTORS
	r.POST("api/v1/investor/signup", handlers.InvestorSignupHandler)
	r.POST("api/v1/investor/login", handlers.InvestorLoginHandler)

	// CORS
	config := cors.DefaultConfig()
  	config.AllowOrigins = []string{"http://127.0.0.1:3000"}
   	r.Use(cors.New(config))

}

func main() {

	// creating gin engine
	r := gin.Default()

	// setting up routes
	setupRoutes(r)

	defer database.DatabaseClient.Disconnect(context.Background())

	// running the server
	r.Run("0.0.0.0:8080")
}
