package services

import (
	"western-movie/cmd/api/models"
	"western-movie/cmd/api/repositories"
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
