package nps

import (
	"context"
	"log"
)

// Tag ...
type Tag struct {
	ID        int     `json:"id"`
	Name      *string `json:"name"`
	Attribute *string `json:"attribute"`
	Number    *int    `json:"number"`
	Timestamp *string `json:"timestamp"`
}

func (r *queryResolver) AllTags(ctx context.Context) ([]*Tag, error) {

	query := `
		SELECT id, name, attribute, number, timestamp
		FROM tags
	`

	rows, err := DB.Queryx(query)
	if err != nil {
		log.Fatalln(err)
	}
	var tags []*Tag
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
