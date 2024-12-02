package songs

import "music-library/internal/models"

func convertRequestUpdateToServ(req request_update, id int64) models.Song {
	return models.Song{
		Id:          id,
		GroupName:   req.Group,
		SongName:    req.Song,
		Text:        req.Text,
		ReleaseDate: req.ReleaseDate.Time,
		Link:        req.Link,
	}
}

func convertResponseGetSong(song models.Song) response_get {
	return response_get{
		Id:          song.Id,
		Group:       song.GroupName,
		Song:        song.SongName,
		ReleaseDate: song.ReleaseDate,
		Text:        song.Text,
		Link:        song.Link,
	}
}
