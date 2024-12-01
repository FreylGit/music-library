package song

import (
	"context"
	modelsServ "music-library/internal/models"
	"time"
)

const limit = 12

func (s *serv) GetSongs(ctx context.Context, page int64, group string, song string, date time.Time) ([]modelsServ.Song, error) {
	offset := (page - 1) * limit
	songs, err := s.songRepo.GetByFilter(ctx, offset, group, song, date)
	if err != nil {
		return nil, err
	}

	return songs, nil
}
