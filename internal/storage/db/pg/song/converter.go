package song

import (
	modelsServ "music-library/internal/models"
	"music-library/internal/storage/db/pg/models"
)

func songRepoToDesc(song models.Song) modelsServ.Song {
	return modelsServ.Song{
		Id:          song.Id,
		GroupName:   song.GroupName,
		SongName:    song.SongName,
		Text:        song.Text,
		ReleaseDate: song.ReleaseDate,
		Link:        song.Link,
	}
}
