package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/topboyasante/mrkt-api/api/v1/routes"
	"github.com/topboyasante/mrkt-api/internal/config"
	"github.com/topboyasante/mrkt-api/internal/database"
)

func main() {
	database.InitDB()
	validator := validator.New()

	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/", func(c *gin.Context) {
		c.String(200, "pong")
	})

	v1 := r.Group("/api/v1")
	{
		routes.AuthRoutes(v1, validator)
		routes.ListingRoutes(v1, validator)
		routes.UserRoutes(v1, validator)
	}

	if err := r.Run(":" + config.ENV.ServerPort); err != nil {
		log.Panicf("error: %s", err)
	}

	log.Printf("server running on port: %s", config.ENV.ServerPort)
}
