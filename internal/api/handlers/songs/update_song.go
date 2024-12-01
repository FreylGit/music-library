package songs

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"time"
)

// UpdateSong Обновление песни
// @Summary Обновляет информацию о песне
// @Description Позволяет обновить информацию о песне по ID. Все поля необязательные, обновляются только переданные значения.
// @Tags Songs
// @Accept json
// @Produce json
// @Param id path int true "ID песни"
// @Param request body request_update false "Информация для обновления песни"
// @Success 200 {string} string "message: Successfully"
// @Failure 400 {string} string "error: Invalid body"
// @Failure 500 {string} string "error: Internal server"
// @Router /songs/{id} [put]
func (s *SongHandler) UpdateSong(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Debug(err.Error())
		c.JSON(http.StatusBadRequest, "Invalid param id")
		return
	}
	var request request_update
	if err := c.BindJSON(&request); err != nil {
		zap.L().Debug(err.Error())
		c.JSON(http.StatusBadRequest, "Invalid body")
		return
	}

	model := requestUpdateToServ(request, id)
	ctx := c.Request.Context()
	err = s.songServ.Update(ctx, model)
	if err != nil {
		zap.L().Debug(err.Error())
		c.JSON(http.StatusInternalServerError, "Internal server")
		return
	}

	c.JSON(http.StatusOK, "Successfully")
}

type request_update struct {
	Group       string     `json:"group,omitempty"`
	Song        string     `json:"song,omitempty"`
	ReleaseDate customTime `json:"release_date,omitempty"`
	Text        string     `json:"text,omitempty"`
	Link        string     `json:"link,omitempty"`
}

type customTime struct {
	time.Time
}

// UnmarshalJSON реализует кастомный парсинг для customTime
func (ct *customTime) UnmarshalJSON(b []byte) error {
	str := string(b)
	str = str[1 : len(str)-1] // Убираем кавычки
	layout := "02.01.2006"    // Указываем ожидаемый формат
	parsedTime, err := time.Parse(layout, str)
	if err != nil {
		return err
	}
	ct.Time = parsedTime
	return nil
}

// MarshalJSON реализует сериализацию customTime обратно в JSON
func (ct customTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + ct.Time.Format("02.01.2006") + `"`), nil
}
