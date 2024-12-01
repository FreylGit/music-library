package songs

import "music-library/internal/models"

func requestAddToServ(req request_add, resp response_add) models.Song {
	return models.Song{
		GroupName:   req.Group,
		SongName:    req.Song,
		Text:        resp.Text,
		ReleaseDate: resp.ReleaseDate,
		Link:        resp.Link,
	}
}

func requestUpdateToServ(req request_update, id int64) models.Song {
	return models.Song{
		Id:          id,
		GroupName:   req.Group,
		SongName:    req.Song,
		Text:        req.Text,
		ReleaseDate: req.ReleaseDate.Time,
		Link:        req.Link,
	}
}

func responseGetSong(song models.Song) response_get {
	return response_get{
		Id:          song.Id,
		Group:       song.GroupName,
		Song:        song.SongName,
		ReleaseDate: song.ReleaseDate,
		Text:        song.Text,
		Link:        song.Link,
	}
}
