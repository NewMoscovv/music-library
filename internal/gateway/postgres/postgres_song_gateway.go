package postgres

import (
	"Music-library/internal/models"
	myLogger "Music-library/pkg/logger"
	"gorm.io/gorm"
)

type PgSongGateway struct {
	db *gorm.DB
}

func (pg *PgSongGateway) GetSongs(filter map[string]string, limit, offset int) ([]models.Song, error) {
	var songs []models.Song
	query := pg.db

	if group, ok := filter["group"]; ok {
		query = query.Where("group = ?", group)
	}
	if song, ok := filter["song"]; ok {
		query = query.Where("song = ?", song)
	}

	// запрашиваем песни
	err := query.Limit(limit).Offset(offset).Find(&songs).Error
	if err != nil {
		myLogger.Error("Ошибка при получении песен", map[string]interface{}{"error": err.Error()})
		return nil, err
	}
	return songs, nil

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
