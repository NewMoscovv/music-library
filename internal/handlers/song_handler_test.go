package handlers_test

import (
	"Music-library/internal/gateway/mocks"
	"Music-library/internal/handlers"
	"Music-library/internal/models"
	myLogger "Music-library/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetSongs(t *testing.T) {
	myLogger.Init("debug")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockSongGateway(ctrl)

	mockRepo.EXPECT().
		GetSongs(gomock.Any(), 10, 0).
		Return([]models.Song{{
			ID: 1, Group: "SHAMAN", Song: "YA RUSSKI"},
		}, nil)

	handler := handlers.NewSongHandler(mockRepo)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/songs", handler.GetSongs)

	req, _ := http.NewRequest(http.MethodGet, "/songs?limit=10&page=1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "SHAMAN")
}

func TestAddSong(t *testing.T) {
	myLogger.Init("debug")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockSongGateway(ctrl)

	input := `{
  "group": "Imagine Dragons",
  "song": "Believer"
 }`

	// Подменяем внешний API
	originalFetch := handlers.FetchSongDetails
	handlers.FetchSongDetails = func(group, song string) (*models.SongDetail, error) {
		return &models.SongDetail{
			ReleaseDate: "2017-02-01",
			Text:        "First things first...",
			Link:        "https://youtube.com/believer",
		}, nil
	}
	defer func() { handlers.FetchSongDetails = originalFetch }()

	// Ожидаем вызов CreateSong
	mockRepo.EXPECT().
		CreateSong(gomock.Any()).
		DoAndReturn(func(song *models.Song) error {
			assert.Equal(t, "Imagine Dragons", song.Group)
			assert.Equal(t, "Believer", song.Song)
			assert.Equal(t, "2017-02-01", song.ReleaseDate)
			return nil
		})

	handler := handlers.NewSongHandler(mockRepo)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/songs", handler.AddSong)

	req := httptest.NewRequest(http.MethodPost, "/songs", strings.NewReader(input))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Believer")
}
