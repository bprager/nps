package nps

import (
	"context"
	"database/sql"
	"log"
)

// Tag ...
type Tag struct {
	ID        int             `json:"id"`
	Name      string          `json:"name"`
	Attribute *sql.NullString `json:"attribute"`
	Number    sql.NullInt64   `json:"number"`
	Timestamp *sql.NullString `json:"timestamp"`
}

func (r *queryResolver) AllTags(ctx context.Context) ([]*Tag, error) {
	var tags []*Tag
	query := `
		SELECT id, name, attribute, number, timestamp
		FROM tags
	`

	rows, err := DB.Queryx(query)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		t := new(Tag)
		err := rows.StructScan(t)
		if err != nil {
			log.Fatalln(err)
		}
		tags = append(tags, t)
	}
	return tags, nil
}

func (r *mutationResolver) AddTag(ctx context.Context, name string, attribute *string, number *int, timestamp *string) (bool, error) {
	query := `
		INSERT INTO tags (name, attribute, number, timestamp)
		VALUES (?1, ?2, ?3, ?4)
	`
	DB.MustExec(query, name, attribute, number, timestamp)

	return true, nil
}
