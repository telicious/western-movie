package services

import (
	"western-movies/internal/core/dto"
	"western-movies/internal/repositories"
)

type MovieService struct {
	repo *repositories.MovieRepository
}

func NewMovieService(repo *repositories.MovieRepository) *MovieService {
	return &MovieService{repo: repo}
}

func (s *MovieService) GetAllMovies() []dto.Movie {
	return s.repo.GetAllMovies()
}

func (s *MovieService) GetMovieByID(id int) (*dto.Movie, error) {
	return s.repo.GetMovieByID(id)
}
