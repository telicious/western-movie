package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"western-movie/cmd/api/repositories"
	"western-movie/cmd/api/services"
	"western-movie/cmd/api/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	f, _ := os.Getwd()
	projectPath, err := filepath.Abs(f)
	if err != nil {
		log.Fatalf("Failed to get absolute path: %v", err)
	}
	jsonPath := path.Join(projectPath, "cmd", "api", "assets/western_movies.json")

	// Load Movies from JSON
	movieRepo, err := repositories.NewMovieRepository(jsonPath)
	if err != nil {
		log.Fatalf("Failed to load movie repository: %v", err)
	}

	// Initialize Service
	movieService := services.NewMovieService(movieRepo)

	// Create a Gin router with default middleware
	router := gin.Default()

	// API Versioning (optional, adjust as needed)
	v1 := router.Group("/api/v1")
	{
		// GET Endpoints
		v1.GET("/movies", func(ctx *gin.Context) {
			movies := movieService.GetAllMovies()
			utils.RespondWithJSON(ctx, http.StatusOK, movies)
		})

		v1.GET("/movies/:id", func(ctx *gin.Context) {
			id := ctx.Param("id")
			intID, err := utils.ParseInt(id)
			if err != nil {
				utils.RespondWithError(ctx, http.StatusBadRequest, "id must be an integer")
				return
			}
			movie, err := movieService.GetMovieByID(intID)
			if err != nil {
				utils.RespondWithError(ctx, http.StatusNotFound, err.Error())
				return
			}
			utils.RespondWithJSON(ctx, http.StatusOK, movie)
		})
	}

	// Start Server
	fmt.Println("Server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
