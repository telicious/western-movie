package repositories

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"western-movies/internal/core/dto"
)

type MovieRepository struct {
	movies []dto.Movie
}

func NewMovieRepository(jsonPath string) (*MovieRepository, error) {
	repo := &MovieRepository{}
	err := repo.loadMovies(jsonPath)
	return repo, err
}

func (r *MovieRepository) loadMovies(jsonPath string) error {
	var data struct {
		WesternMovies []dto.Movie `json:"WesternMovies"`
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

func (r *MovieRepository) GetAllMovies() []dto.Movie {
	return r.movies
}

func (r *MovieRepository) GetMovieByID(id int) (*dto.Movie, error) {
	for _, movie := range r.movies {
		if movie.ID == id {
			return &movie, nil
		}
	}
	return nil, fmt.Errorf("movie with ID %d not found", id)
}
