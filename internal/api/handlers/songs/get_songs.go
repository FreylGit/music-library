package songs

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"time"
)

// GetSongs Получение списка песен
// @Summary Получает список песен с фильтрацией
// @Description Позволяет получить список песен с учетом фильтров и пагинации
// @Tags Songs
// @Accept json
// @Produce json
// @Param page query int true "Номер страницы для пагинации (начиная с 1)"
// @Param group query string false "Название группы для фильтрации"
// @Param song query string false "Название песни для фильтрации"
// @Param releaseDate query string false "Дата релиза песни в формате DD.MM.YYYY"
// @Success 200 {object} response_songs "Список песен"
// @Failure 400 {string} string "error: Invalid param id"
// @Failure 500 {string} string "error: Internal server"
// @Router /songs [get]
func (s *SongHandler) GetSongs(c *gin.Context) {
	pageStr := c.Query("page")
	page, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		zap.L().Debug(err.Error())
		c.JSON(http.StatusBadRequest, "Invalid param id")
		return
	}
	group := c.Query("group")
	song := c.Query("song")
	releaseDateStr := c.Query("releaseDate")
	const layout = "02.01.2006"
	date, _ := time.Parse(layout, releaseDateStr)

	ctx := c.Request.Context()
	songs, err := s.songServ.GetSongs(ctx, page, group, song, date)

	responseDate := response_songs{
		Items: make([]response_get, len(songs)),
	}
	for i := 0; i < len(responseDate.Items); i++ {
		responseDate.Items[i] = convertResponseGetSong(songs[i])
	}
	c.JSON(http.StatusOK, responseDate)
}

type response_songs struct {
	Items []response_get `json:"items"`
}
