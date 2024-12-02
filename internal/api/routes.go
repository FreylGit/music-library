package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"music-library/docs"
	"music-library/internal/api/handlers/songs"
)

func RegisterRoutes(router *gin.Engine, songHandler *songs.SongHandler) {
	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Группа маршрутов для песен
	songGroup := router.Group("/songs")
	{
		songGroup.GET("/", songHandler.GetSongs)         // Список песен
		songGroup.GET("/:id", songHandler.GetSong)       // Получение песни по Id
		songGroup.POST("/", songHandler.AddSong)         // Добавление новой песни
		songGroup.PUT("/:id", songHandler.UpdateSong)    // Изменение песни
		songGroup.DELETE("/:id", songHandler.DeleteSong) // Удаление песни
	}
}
