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

	if group, ok := filter["group"]; ok && group != "" {
		query = query.Where("\"group\" = ?", group)
	}
	if song, ok := filter["song"]; ok && song != "" {
		query = query.Where("song = ?", song)
	}

	// запрашиваем песни
	err := query.Limit(limit).Offset(offset).Find(&songs).Error
	if err != nil {
		myLogger.Error("Ошибка при получении песен", map[string]interface{}{"error": err.Error()})
		return nil, err
	}
	myLogger.Debug("Песни из БД", map[string]interface{}{"count": len(songs)})
	return songs, nil

}

func (pg *PgSongGateway) GetSongByID(id uint) (*models.Song, error) {
	var song models.Song
	err := pg.db.First(&song, id).Error
	if err != nil {
		myLogger.Error("Ошибка при получении песни", map[string]interface{}{"error": err.Error()})
		return nil, err
	}
	return &song, nil
}

func (pg *PgSongGateway) CreateSong(song *models.Song) error {
	err := pg.db.Create(song).Error
	if err != nil {
		myLogger.Error("Ошибка при добавлении песни", map[string]interface{}{"error": err.Error()})
	}
	return err
}

func (pg *PgSongGateway) UpdateSong(song *models.Song) error {
	err := pg.db.Save(song).Error
	if err != nil {
		myLogger.Error("Ошибка при обновлении песни", map[string]interface{}{"error": err.Error()})
	}
	return err
}

func (pg *PgSongGateway) DeleteSong(id uint) error {
	err := pg.db.Delete(&models.Song{}, id).Error
	if err != nil {
		myLogger.Error("Ошибка при удалении песни", map[string]interface{}{"error": err.Error()})
	}
	return err
}

func NewPgSongGateway(db *gorm.DB) *PgSongGateway {
	return &PgSongGateway{db: db}
}
