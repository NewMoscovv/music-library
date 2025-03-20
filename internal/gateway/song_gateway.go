package gateway

import "Music-library/internal/models"

type SongGateway interface {
	GetSongs(filter map[string]string, limit, offset int) ([]models.Song, error)
	GetSongByID(id uint) (*models.Song, error)
	CreateSong(song *models.Song) error
	UpdateSong(song *models.Song) error
	DeleteSong(id uint) error
}
