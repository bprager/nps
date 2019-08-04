package nps

import (
	"context"
	"log"
)

// User ...
type User struct {
	ID         int         `json:"id"`
	Email      *string     `json:"email"`
	FirstName  *string     `json:"firstName" db:"first_name"`
	LastName   *string     `json:"lastName" db:"last_name"`
	NickName   *string     `json:"nickName" db:"nick_name"`
	Orgs       []*Org      `json:"orgs" db:"orgs"`
	Tags       []*Tag      `json:"tags"`
	Categories []*Category `json:"categories"`
}

func (r *queryResolver) AllUsers(ctx context.Context) ([]*User, error) {

	query := `
		SELECT id, first_name, last_name, nick_name
		FROM users
	`

	rows, err := DB.Queryx(query)
	if err != nil {
		log.Fatalln(err)
	}
	var users []*User
	for rows.Next() {
		u := new(User)
		err := rows.StructScan(u)
		if err != nil {
			log.Fatalln(err)
		}
		users = append(users, u)
	}
	return users, nil
}
