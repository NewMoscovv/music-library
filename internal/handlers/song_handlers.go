package handlers

import (
	"Music-library/internal/gateway"
	"Music-library/internal/models"
	myLogger "Music-library/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type SongHandler struct {
	gateway gateway.SongGateway
}

func NewSongHandler(gateway gateway.SongGateway) *SongHandler {
	return &SongHandler{gateway: gateway}
}

func (h *SongHandler) GetSongs(c *gin.Context) {
	filter := map[string]string{
		"group": c.Query("group"),
		"song":  c.Query("song"),
	}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	songs, err := h.gateway.GetSongs(filter, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении песен"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"total": len(songs), "data": songs})

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
