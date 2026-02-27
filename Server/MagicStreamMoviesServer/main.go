package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lva100/movie-streaming-service/Server/MagicStreamMoviesServer/controllers"
	"github.com/lva100/movie-streaming-service/Server/MagicStreamMoviesServer/database"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func main() {
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, MagicStreamMovies!")
	})

	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: unable to find .env file")
	}

	var client *mongo.Client = database.Connect()

	if err := client.Ping(context.Background(), nil); err != nil {
		log.Fatalf("Failed to reach server: %v", err)
	}
	defer func() {
		err := client.Disconnect(context.Background())
		if err != nil {
			log.Fatalf("Failed to disconnect from MongoDB: %v", err)
		}

	}()

	router.GET("/movies", controllers.GetMovies(client))
	router.GET("/movie/:imdb_id", controllers.GetMovie(client))
	router.POST("/addmovie", controllers.AddMovie(client))
	router.POST("/register", controllers.RegisterUser(client))
	router.POST("/login", controllers.LoginUser(client))

	if err := router.Run(":8088"); err != nil {
		fmt.Println("Failed to start server", err)
	}
}
