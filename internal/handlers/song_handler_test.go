package handlers_test

import (
	"Music-library/internal/gateway/mocks"
	"Music-library/internal/handlers"
	"Music-library/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetSongs(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockSongGateway(ctrl)

	mockRepo.EXPECT().
		GetSongs(gomock.Any(), 10, 0).
		Return([]models.Song{{
			Group: "SHAMAN", Song: "YA RUSSKI"},
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
