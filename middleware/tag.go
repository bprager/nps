package nps

import (
	"context"
	"log"
)

// Tag ...
type Tag struct {
	ID        int         `json:"id"`
	Name      string      `json:"name"`
	Attribute *NullString `json:"attribute"`
	Number    NullInt64   `json:"number"`
	Timestamp *NullString `json:"timestamp"`
}

// AllTags ...
func (r *queryResolver) AllTags(ctx context.Context, limit int, offset int) (*TagsResult, error) {
	result := new(TagsResult)
	var tags []*Tag
	query := `
		SELECT id, name, attribute, number, timestamp, count(*)
		FROM tags
	`

	rows, err := DB.Queryx(query)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		t := new(Tag)
		err := rows.Scan(&t.ID, &t.Name, &t.Attribute, &t.Number, &t.Timestamp, &result.TotalCount)
		if err != nil {
			log.Fatalln(err)
		}
		tags = append(tags, t)
	}
	result.Tags = tags
	return result, nil
}

// AddTag ...
func (r *mutationResolver) AddTag(ctx context.Context, name string, attribute *string, number *int, timestamp *string) (bool, error) {
	query := `
		INSERT INTO tags (name, attribute, number, timestamp)
		VALUES (?1, ?2, ?3, ?4)
	`
	DB.MustExec(query, name, attribute, number, timestamp)

	return true, nil
}
