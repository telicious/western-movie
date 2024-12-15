package services

import (
	"western-movies/cmd/api/models"
	"western-movies/cmd/api/repositories"
)

type MovieService struct {
	repo *repositories.MovieRepository
}

func NewMovieService(repo *repositories.MovieRepository) *MovieService {
	return &MovieService{repo: repo}
}

func (s *MovieService) GetAllMovies() []models.Movie {
	return s.repo.GetAllMovies()
}

func (s *MovieService) GetMovieByID(id int) (*models.Movie, error) {
	return s.repo.GetMovieByID(id)
}
