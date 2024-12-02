package song

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	modelsServ "music-library/internal/models"
	"net/http"
	"time"
)

func (s *serv) Add(ctx context.Context, song modelsServ.Song) error {
	// TODO: если я не так понял задание, можно проверить, заменив на debug_request_info
	info, err := request_info(song.GroupName, song.SongName)
	if err != nil {
		return err
	}
	song.Text = info.Text
	song.ReleaseDate = info.ReleaseDate
	song.Link = info.Link
	err = s.songRepo.Add(ctx, song)
	if err != nil {
		return err
	}
	return nil
}

func request_info(group string, song string) (response_add, error) {
	// Тут поненять на нужный хост
	const host = "http://localhost:8080"
	resp, err := http.Get(fmt.Sprintf("%s/info?group=%s&song=%s", host, group, song))
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

type response_add struct {
	ReleaseDate time.Time `json:"releaseDate"`
	Text        string    `json:"text"`
	Link        string    `json:"link"`
}
