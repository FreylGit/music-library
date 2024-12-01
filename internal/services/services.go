package services

import (
	"context"
	"music-library/internal/models"
	"time"
)

type SongService interface {
	Add(ctx context.Context, song models.Song) error
	Delete(ctx context.Context, id int64) error
	GetSong(ctx context.Context, id int64, kuplet int64) (models.Song, error)
	GetSongs(ctx context.Context, page int64, group string, song string, date time.Time) ([]models.Song, error)
	Update(ctx context.Context, song models.Song) error
}
