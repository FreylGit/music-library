package songs

import (
	"music-library/internal/services"
)

type SongHandler struct {
	songServ services.SongService
}

func NewSongHandler(songServ services.SongService) *SongHandler {
	return &SongHandler{songServ: songServ}
}
