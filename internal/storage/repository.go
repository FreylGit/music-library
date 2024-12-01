package storage

import (
	"context"
	modelsServ "music-library/internal/models"
	"time"
)

type SongRepository interface {
	GetByFilter(ctx context.Context, offset int64, group string, song string, releaseDate time.Time) ([]modelsServ.Song, error)
	Get(ctx context.Context, id int64) (modelsServ.Song, error)
	Delete(ctx context.Context, id int64) error
	Edit(ctx context.Context, song modelsServ.Song) error
	Add(ctx context.Context, song modelsServ.Song) error
}
