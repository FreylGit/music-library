package song

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	modelsServ "music-library/internal/models"
	"music-library/internal/storage"
	"music-library/internal/storage/db/pg/models"
	"time"
)

const (
	tableName         = "songs"
	idColumn          = "id"
	groupNameColumn   = "group_name"
	songNameColumn    = "song_name"
	releaseDateColumn = "release_date"
	textColumn        = "text"
	linkColumn        = "link"
)

type repo struct {
	db *pgxpool.Pool
}

func NewSongRepository(db *pgxpool.Pool) storage.SongRepository {
	return &repo{
		db: db,
	}
}

// Получение данных библиотеки с фильтрацией по всем полям
func (r *repo) GetByFilter(ctx context.Context, offset int64, group string, song string, releaseDate time.Time) ([]modelsServ.Song, error) {
	const limit = 12

	builder := squirrel.Select(idColumn, groupNameColumn, songNameColumn, releaseDateColumn, textColumn, linkColumn).
		From(tableName).
		Limit(uint64(limit)).
		Offset(uint64(offset)).
		PlaceholderFormat(squirrel.Dollar)

	// Добавляем фильтры только для заданных параметров
	if group != "" {
		builder = builder.Where(squirrel.Expr("LOWER("+groupNameColumn+") = LOWER(?)", group))
	}
	if song != "" {
		builder = builder.Where(squirrel.Expr("LOWER("+songNameColumn+") = LOWER(?)", song))
	}
	if !releaseDate.IsZero() {
		builder = builder.Where(squirrel.Eq{releaseDateColumn: releaseDate})
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var songs []modelsServ.Song
	for rows.Next() {
		var song models.Song

		if err := rows.Scan(
			&song.Id,
			&song.GroupName,
			&song.SongName,
			&song.ReleaseDate,
			&song.Text,
			&song.Link,
		); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		songs = append(songs, songRepoToDesc(song))
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return songs, nil
}

// Получение текста песни
func (r *repo) Get(ctx context.Context, id int64) (modelsServ.Song, error) {
	builder := squirrel.
		Select(idColumn, groupNameColumn, songNameColumn, releaseDateColumn, textColumn, linkColumn).
		From(tableName).
		Where(squirrel.Eq{idColumn: id}).
		PlaceholderFormat(squirrel.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return modelsServ.Song{}, err
	}
	row := r.db.QueryRow(ctx, query, args...)
	var model models.Song
	err = row.Scan(&model.Id, &model.GroupName, &model.SongName, &model.ReleaseDate, &model.Text, &model.Link)
	if err != nil {
		return modelsServ.Song{}, err
	}
	return songRepoToDesc(model), nil
}

// Удаление песни
func (r *repo) Delete(ctx context.Context, id int64) error {
	builder := squirrel.Delete(tableName).
		Where(squirrel.Eq{idColumn: id}).
		PlaceholderFormat(squirrel.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}
	tag, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return fmt.Errorf("no rows found to delete")
	}
	return nil
}

// Изменение данных песни
func (r *repo) Edit(ctx context.Context, song modelsServ.Song) error {
	builder := squirrel.Update(tableName).
		Where(squirrel.Eq{idColumn: song.Id})

	if song.GroupName != "" {
		builder = builder.Set(groupNameColumn, song.GroupName)
	}
	if song.SongName != "" {
		builder = builder.Set(songNameColumn, song.SongName)
	}
	if song.Text != "" {
		builder = builder.Set(textColumn, song.Text)
	}
	if song.Link != "" {
		builder = builder.Set(linkColumn, song.Link)
	}
	if !song.ReleaseDate.IsZero() {
		builder = builder.Set(releaseDateColumn, song.ReleaseDate)
	}

	query, args, err := builder.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return err
	}

	tag, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return fmt.Errorf("no rows updated")
	}
	return nil
}

// Добавление новой песни
func (r *repo) Add(ctx context.Context, song modelsServ.Song) error {
	builder := squirrel.Insert(tableName).
		Columns(groupNameColumn, songNameColumn, textColumn, releaseDateColumn, linkColumn).
		Values(
			song.GroupName,
			song.SongName,
			song.Text,
			song.ReleaseDate,
			song.Link,
		).
		PlaceholderFormat(squirrel.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	tag, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return fmt.Errorf("Failed to write song to table")
	}

	return nil
}
