package handlers

import (
	"Music-library/internal/gateway"
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
