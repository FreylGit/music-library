package models

import "time"

type Song struct {
	Id          int64
	GroupName   string
	SongName    string
	ReleaseDate time.Time
	Text        string
	Link        string
}
