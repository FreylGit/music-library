package songs

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
	"music-library/internal/models"
	"net/http"
	"time"
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

func request_info(group string, song string) (response_add, error) {
	resp, err := http.Get(fmt.Sprintf("/info?group=%s&song=%s", group, song))
	if err != nil {
		return response_add{}, fmt.Errorf("External service error")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response_add{}, fmt.Errorf("Error reading response body")
	}
	var response response_add
	if err := json.Unmarshal(body, &response); err != nil {
		return response_add{}, fmt.Errorf("Error parsing response body")
	}

	return response, nil
}

func debug_request_info(group string, song string) (response_add, error) {
	const layout = "02.01.2006"
	date, _ := time.Parse(layout, "16.07.2006")
	return response_add{
		ReleaseDate: date,
		Link:        "https://www.youtube.com/watch?v=Xsp3_a-PMTw",
		Text:        " Ooh baby, don't you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses\n How long before you let me go?\n\nOoh\nYou set my soul alight\nOoh\nYou set my soul alight",
	}, nil
}

// request_add - структура для данных запроса при добавлении песни
// @Description Структура данных для добавления песни
type request_add struct {
	Group string `json:"group"` // Группа, обязательное поле
	Song  string `json:"song"`  // Песня, обязательное поле
}

type response_add struct {
	ReleaseDate time.Time `json:"releaseDate"`
	Text        string    `json:"text"`
	Link        string    `json:"link"`
}
