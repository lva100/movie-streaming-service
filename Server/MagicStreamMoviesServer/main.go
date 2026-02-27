package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lva100/movie-streaming-service/Server/MagicStreamMoviesServer/controllers"
)

func main() {
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, MagicStreamMovies!")
	})

	router.GET("/movies", controllers.GetMovies())
	router.GET("/movie/:imdb_id", controllers.GetMovie())
	router.POST("/addmovie", controllers.AddMovie())

	if err := router.Run(":8088"); err != nil {
		fmt.Println("Failed to start server", err)
	}
}
