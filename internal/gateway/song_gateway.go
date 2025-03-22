package gateway

import "Music-library/internal/models"

//go:generate mockgen -source=song_gateway.go -destination=mocks/mock_gateway.go -package=mocks
type SongGateway interface {
	GetSongs(filter map[string]string, limit, offset int) ([]models.Song, error)
	GetSongByID(id uint) (*models.Song, error)
	CreateSong(song *models.Song) error
	UpdateSong(song *models.Song) error
	DeleteSong(id uint) error
}
