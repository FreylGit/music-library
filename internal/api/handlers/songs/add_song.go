package songs

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"music-library/internal/models"
	"net/http"
)

// AddSong добавляет новую песню в библиотеку
// @Summary Добавление новой песни
// @Description Позволяет добавить песню в библиотеку, принимая название группы и песни
// @Tags Songs
// @Accept json
// @Produce json
// @Param request body request_add true "Данные для добавления песни"
// @Success 200 {string} string "message: Successfully"
// @Failure 400 {string} string "error: Invalid body"
// @Failure 500 {string} string "error: Internal server"
// @Router /songs [post]
func (s *SongHandler) AddSong(c *gin.Context) {
	var request request_add

	if err := c.BindJSON(&request); err != nil {
		zap.L().Debug(err.Error())
		c.JSON(http.StatusBadRequest, "Invalid body")
		return
	}
	ctx := c.Request.Context()

	// Конвертим то что приходит к нашей бизнес логике
	model := models.Song{
		GroupName: request.Group,
		SongName:  request.Song,
	}
	err := s.songServ.Add(ctx, model)

	if err != nil {
		zap.L().Debug(err.Error())
		c.JSON(http.StatusInternalServerError, "Internal server")
		return
	}
	c.JSON(http.StatusOK, "Successfully")
}

// request_add - структура для данных запроса при добавлении песни
// @Description Структура данных для добавления песни
type request_add struct {
	Group string `json:"group"` // Группа, обязательное поле
	Song  string `json:"song"`  // Песня, обязательное поле
}
