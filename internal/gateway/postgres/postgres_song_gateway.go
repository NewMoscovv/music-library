package postgres

import (
	"Music-library/internal/models"
	"gorm.io/gorm"
)

type PgSongGateway struct {
	db *gorm.DB
}

func (pg *PgSongGateway) GetSongs(filter map[string]string, limit, offset int) ([]models.Song, error) {
	//TODO implement me
	panic("implement me")
}

func (pg *PgSongGateway) GetSongByID(id int) (*models.Song, error) {
	//TODO implement me
	panic("implement me")
}

func (pg *PgSongGateway) CreateSong(song *models.Song) error {
	//TODO implement me
	panic("implement me")
}

func (pg *PgSongGateway) UpdateSong(song *models.Song) error {
	//TODO implement me
	panic("implement me")
}

func (pg *PgSongGateway) DeleteSong(id int) error {
	//TODO implement me
	panic("implement me")
}

func NewPgSongGateway(db *gorm.DB) *PgSongGateway {
	return &PgSongGateway{db: db}
}
