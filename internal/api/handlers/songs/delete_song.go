package songs

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

// DeleteSong удаляет песню по указанному ID
// @Summary Удаление песни по ID
// @Description Позволяет удалить песню из библиотеки по её ID
// @Tags Songs
// @Accept json
// @Produce json
// @Param id path int64 true "ID песни, которую нужно удалить"
// @Success 200 {string} string "message: Successfully"
// @Failure 400 {string} string "error: Invalid param id"
// @Failure 500 {string} string "error: Internal server"
// @Router /songs/{id} [delete]
func (s *SongHandler) DeleteSong(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {
		zap.L().Debug(err.Error())
		c.JSON(http.StatusBadRequest, "Invalid param id")
		return
	}
	ctx := context.Background()
	err = s.songServ.Delete(ctx, id)
	if err != nil {
		zap.L().Debug(err.Error())
		c.JSON(http.StatusInternalServerError, "Internal server")
		return
	}
	c.JSON(http.StatusOK, "Successfully")
}
