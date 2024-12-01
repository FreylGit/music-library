package song

import (
	"context"
	modelsServ "music-library/internal/models"
	"strings"
)

func (s *serv) GetSong(ctx context.Context, id int64, kuplet int64) (modelsServ.Song, error) {
	song, err := s.songRepo.Get(ctx, id)
	if err != nil {
		return modelsServ.Song{}, err
	}
	if kuplet != 0 {
		song.Text = filterTextByKuplet(song.Text, kuplet)
	}
	return song, nil
}

// Пагинации текста песни по куплету
func filterTextByKuplet(text string, kuplet int64) string {
	kuplets := strings.Split(text, "\n\n")
	if kuplet > 0 && kuplet <= int64(len(kuplets)) {
		return kuplets[kuplet-1]
	}
	// Возвращаем полный текст если куплет не найден
	return text
}
