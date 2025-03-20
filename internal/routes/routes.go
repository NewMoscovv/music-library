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
		songRoutes.GET("/:id", songHandler.GetSongByID)
		songRoutes.POST("", songHandler.AddSong)
		songRoutes.PUT("/:id", songHandler.UpdateSong)
		songRoutes.DELETE("/:id", songHandler.DeleteSong)
		songRoutes.GET("/:id/lyrics", songHandler.GetLyrics)
	}

}
