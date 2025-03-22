package handlers

import (
	"Music-library/internal/gateway"
	"Music-library/internal/models"
	myLogger "Music-library/pkg/logger"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type SongHandler struct {
	gateway gateway.SongGateway
}

func NewSongHandler(gateway gateway.SongGateway) *SongHandler {
	return &SongHandler{gateway: gateway}
}

// GetSongs получает список песен с фильтрацией и пагинацией
// @Summary Получение песен
// @Description Получение песен с фильтрацией по всем полям и пагинацией
// @Tags Songs
// @Accept json
// @Produce json
// @Param group query string false "Группа"
// @Param song query string false "Песня"
// @Param release_date query string false "Дата релиза"
// @Param text query string false "Часть текста песни"
// @Param link query string false "Ссылка"
// @Param page query int false "Номер страницы"
// @Param limit query int false "Количество записей на странице"
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]string
// @Router /songs [get]
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

// GetSongByID получает песню по ID
// @Summary Получение песни по ID
// @Tags Songs
// @Accept json
// @Produce json
// @Param id path int true "ID песни"
// @Success 200 {object} models.Song
// @Failure 404 {object} map[string]string
// @Router /songs/{id} [get]
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

// AddSong добавляет новую песню и обогащает её через внешний API
// @Summary Добавление песни
// @Tags Songs
// @Accept json
// @Produce json
// @Param song body models.Song true "Песня (group + song)"
// @Success 200 {object} models.Song
// @Failure 400 {object} map[string]string
// @Failure 502 {object} map[string]string
// @Router /songs [post]
func (h *SongHandler) AddSong(c *gin.Context) {
	var input models.Song

	// Получаем group и song из JSON
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Формируем URL внешнего API
	apiUrl := fmt.Sprintf(
		"%s/info?group=%s&song=%s",
		os.Getenv("API_URL"),
		url.QueryEscape(input.Group),
		url.QueryEscape(input.Song),
	)

	// Отправляем GET-запрос
	resp, err := http.Get(apiUrl)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "Не удалось подключиться к внешнему API"})
		return
	}
	defer resp.Body.Close()

	// Проверяем статус ответа
	if resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusBadRequest, gin.H{"error": "API вернул ошибку"})
		return
	}

	// Распарсим тело ответа
	var details models.SongDetail
	if err := json.NewDecoder(resp.Body).Decode(&details); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка обработки ответа API"})
		return
	}

	// Обогащаем и сохраняем песню
	song := models.Song{
		Group:       input.Group,
		Song:        input.Song,
		ReleaseDate: details.ReleaseDate,
		Text:        details.Text,
		Link:        details.Link,
	}

	if err := h.gateway.CreateSong(&song); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сохранения песни"})
		return
	}

	c.JSON(http.StatusOK, song)
}

// UpdateSong обновляет данные песни по ID
// @Summary Обновление песни
// @Tags Songs
// @Accept json
// @Produce json
// @Param id path int true "ID песни"
// @Param song body models.Song true "Обновлённые данные"
// @Success 200 {object} models.Song
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /songs/{id} [put]
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

// DeleteSong удаляет песню по ID
// @Summary Удаление песни
// @Tags Songs
// @Accept json
// @Produce json
// @Param id path int true "ID песни"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /songs/{id} [delete]
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

// GetLyrics возвращает текст песни постранично (по куплетам)
// @Summary Получение текста песни с пагинацией
// @Tags Songs
// @Accept json
// @Produce json
// @Param id path int true "ID песни"
// @Param page query int false "Номер страницы (куплет)"
// @Param limit query int false "Количество куплетов"
// @Success 200 {array} string
// @Failure 404 {object} map[string]string
// @Router /songs/{id}/lyrics [get]
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
