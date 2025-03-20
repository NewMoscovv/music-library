package routes

import (
	"Music-library/internal/gateway"
	"Music-library/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, songGateway gateway.SongGateway) {
	songHandler := handlers.NewSongHandler(songGateway)

	songRoutes := router.Group("/songs")
	{
		songRoutes.GET("", songHandler.GetSongs)
	}

}
