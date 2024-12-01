package song

import (
	"music-library/internal/services"
	"music-library/internal/storage"
)

type serv struct {
	songRepo storage.SongRepository
}

func NewSongService(songRepo storage.SongRepository) services.SongService {
	return &serv{songRepo: songRepo}
}
