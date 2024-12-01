package song

import (
	"context"
	modelsServ "music-library/internal/models"
)

func (s *serv) Update(ctx context.Context, song modelsServ.Song) error {
	err := s.songRepo.Edit(ctx, song)
	if err != nil {
		return err
	}

	return nil
}
