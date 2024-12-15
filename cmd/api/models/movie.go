package models

type Rating struct {
	IMDB           float64 `json:"imdb"`
	RottenTomatoes int     `json:"rottenTomatoes"`
}

type Movie struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Director    string   `json:"director"`
	ReleaseYear int      `json:"releaseYear"`
	Starring    []string `json:"starring"`
	Genre       string   `json:"genre"`
	Awards      []string `json:"awards"`
	Rating      Rating   `json:"rating"`
	Synopsis    string   `json:"synopsis"`
}
