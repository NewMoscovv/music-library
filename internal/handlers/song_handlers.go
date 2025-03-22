package handlers

import (
	"Music-library/internal/gateway"
	"Music-library/internal/models"
	myLogger "Music-library/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type SongHandler struct {
	gateway gateway.SongGateway
}

func NewSongHandler(gateway gateway.SongGateway) *SongHandler {
	return &SongHandler{gateway: gateway}
}

func (h *SongHandler) GetSongs(c *gin.Context) {
	filter := map[string]string{
		"group":        c.Query("group"),
		"song":         c.Query("song"),
		"release_date": c.Query("release_date"),
		"text":         c.Query("text"),
		"link":         c.Query("link"),
	}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	songs, err := h.gateway.GetSongs(filter, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении песен"})
		return
	}

	myLogger.Debug("Отправляем песни клиенту", map[string]interface{}{"total1": len(songs)})

	c.JSON(http.StatusOK, gin.H{"total": len(songs), "data": songs})

}

func (h *SongHandler) GetSongByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID"})
		return
	}
	song, err := h.gateway.GetSongByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Песня не найдена"})
		return
	}
	c.JSON(http.StatusOK, song)
}

func (h *SongHandler) AddSong(c *gin.Context) {
	var song models.Song

	if err := c.ShouldBindJSON(&song); err != nil {
		myLogger.Error("Ошибка валидации данных", map[string]interface{}{"error": err.Error()})
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.gateway.CreateSong(&song); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	myLogger.Info("Песня успешно добавлена", map[string]interface{}{"song": song})
	c.JSON(http.StatusOK, song)
}

func (h *SongHandler) UpdateSong(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID"})
		return
	}

	var song models.Song
	if err = c.ShouldBindJSON(&song); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	song.ID = uint(id)
	if err := h.gateway.UpdateSong(&song); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении песни"})
		return
	}

	c.JSON(http.StatusOK, song)
}

func (h *SongHandler) DeleteSong(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID"})
		return
	}

	if err = h.gateway.DeleteSong(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении песни"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Песня удалена"})
}

func (h *SongHandler) GetLyrics(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID"})
		return
	}

	song, err := h.gateway.GetSongByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Песня не найдна"})
		return
	}

	verses := strings.Split(song.Text, "\n\n")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	if offset > len(verses) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Куплет не найден"})
		return
	}

	end := offset + limit
	if end > len(verses) {
		end = len(verses)
	}
	c.JSON(http.StatusOK, verses[offset:end])
}
