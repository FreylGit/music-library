package songs

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"time"
)

// GetSong Получение песни
// @Summary Получает песню
// @Description Позволяет получить песню по id
// @Tags Songs
// @Accept json
// @Produce json
// @Param id path int true "Id песни"
// @Param kuplet query int false "Номер куплета для пагинации"
// @Success 200 {object} response_get "Детали песни"
// @Failure 400 {string} string "error: Invalid param id"
// @Failure 500 {string} string "error: Internal server"
// @Router /songs/{id} [get]
func (s *SongHandler) GetSong(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Debug(err.Error())
		c.JSON(http.StatusBadRequest, "Invalid param id")
		return
	}
	ctx := c.Request.Context()

	kupletStr := c.Query("kuplet")
	kuplet, err := strconv.ParseInt(kupletStr, 10, 64)
	if err != nil {
		zap.L().Debug(err.Error())
		c.JSON(http.StatusBadRequest, "Invalid query param kuplet")
		return
	}
	song, err := s.songServ.GetSong(ctx, id, kuplet)
	if err != nil {
		zap.L().Debug(err.Error())
		c.JSON(http.StatusInternalServerError, "Internal server")
		return
	}
	response := convertResponseGetSong(song)

	c.JSON(http.StatusOK, response)

}

type response_get struct {
	Id          int64     `json:"id"`
	Group       string    `json:"group"`
	Song        string    `json:"song"`
	ReleaseDate time.Time `json:"release_date"`
	Text        string    `json:"text"`
	Link        string    `json:"link"`
}
