package app

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"music-library/internal/api/handlers/songs"
	"music-library/internal/config"
	"music-library/internal/config/env"
	"music-library/internal/services"
	song2 "music-library/internal/services/song"
	"music-library/internal/storage"
	"music-library/internal/storage/db/pg/song"
)

type service_provider struct {
	songRepo    storage.SongRepository
	configHttp  config.ConfigHTTP
	configPg    config.ConfigPG
	db          *pgxpool.Pool
	songHandler *songs.SongHandler
	songServ    services.SongService
}

func NewServiceProvider() *service_provider {
	return &service_provider{}
}

func (sp *service_provider) ConfigHTTP() config.ConfigHTTP {
	if sp.configHttp == nil {
		sp.configHttp = env.NewConfig()
	}
	return sp.configHttp
}

func (sp *service_provider) ConfigPG() config.ConfigPG {
	if sp.configPg == nil {
		sp.configPg = env.NewConfigPG()
	}
	return sp.configPg
}

func (sp *service_provider) SongRepository(ctx context.Context) storage.SongRepository {
	if sp.songRepo == nil {
		db := sp.DB(ctx)
		sp.songRepo = song.NewSongRepository(db)
	}

	return sp.songRepo
}

func (sp *service_provider) DB(ctx context.Context) *pgxpool.Pool {
	if sp.db == nil {
		db, err := pgxpool.New(ctx, sp.ConfigPG().DSN())
		if err != nil {
			log.Fatalf("Failed to connect to DB: %v", err)
		}
		err = db.Ping(ctx)
		if err != nil {
			log.Fatalf("Failed ping DB: %v", err)
		}
		sp.db = db
	}
	return sp.db
}

func (sp *service_provider) Close() {
	if sp.db != nil {
		sp.db.Close()

	}
}

func (sp *service_provider) SongHandler(ctx context.Context) *songs.SongHandler {
	if sp.songHandler == nil {
		sp.songHandler = songs.NewSongHandler(sp.SongService(ctx))
	}

	return sp.songHandler
}

func (sp *service_provider) SongService(ctx context.Context) services.SongService {
	if sp.songServ == nil {
		sp.songServ = song2.NewSongService(sp.SongRepository(ctx))
	}

	return sp.songServ
}
