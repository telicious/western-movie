package repositories

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"western-movies/cmd/api/models"
)

type MovieRepository struct {
	movies []models.Movie
}

func NewMovieRepository(jsonPath string) (*MovieRepository, error) {
	repo := &MovieRepository{}
	err := repo.loadMovies(jsonPath)
	return repo, err
}

func (r *MovieRepository) loadMovies(jsonPath string) error {
	var data struct {
		WesternMovies []models.Movie `json:"WesternMovies"`
	}
	file, err := os.Open(jsonPath)
	if err != nil {
		return err
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		return err
	}

	r.movies = data.WesternMovies
	return nil
}

func (r *MovieRepository) GetAllMovies() []models.Movie {
	return r.movies
}

func (r *MovieRepository) GetMovieByID(id int) (*models.Movie, error) {
	for _, movie := range r.movies {
		if movie.ID == id {
			return &movie, nil
		}
	}
	return nil, fmt.Errorf("movie with ID %d not found", id)
}
