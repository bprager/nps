package nps

import (
	"context"
	"database/sql"
	"log"
)

// NullInt64 is an alias for sql.NullInt64 data type
type NullInt64 struct {
	sql.NullInt64
}

// NullString is an alias for sql.NullInt64 data type
type NullString struct {
	sql.NullString
}

// User ...
type User struct {
	ID         int         `json:"id"`
	Email      NullString  `json:"email"`
	FirstName  NullString  `json:"firstName" db:"first_name"`
	LastName   NullString  `json:"lastName" db:"last_name"`
	NickName   NullString  `json:"nickName" db:"nick_name"`
	Orgs       *[]Org      `json:"orgs" db:"orgs"`
	Tags       *[]Tag      `json:"tags"`
	Categories *[]Category `json:"categories"`
}

// AllUsers ...
func (r *queryResolver) AllUsers(ctx context.Context, limit int, offset int) (*UsersResult, error) {

	query := `
		SELECT id, first_name, last_name, nick_name, orgs, tags, categories, count(*)
		FROM users
	`
	if limit != 0 {
		query += ` LIMIT ? OFFSET ?`
	}

	rows, err := DB.Queryx(query, limit, offset)
	if err != nil {
		log.Fatalln(err)
	}
	result := new(UsersResult)
	var users []*User
	// var orgs []*Org
	// var tags []*Tag
	// var cats []*Category

	count := 0
	for rows.Next() {
		count++
		u := new(User)
		o := new(NullInt64)
		t := new(NullInt64)
		c := new(NullInt64)
		err := rows.Scan(&u.ID, &u.Email, &u.FirstName, &u.LastName, &u.NickName, o, t, c, &result.TotalCount)
		if err != nil {
			log.Fatalln(err)
		}
		if o.Valid || t.Valid || c.Valid {
			// query := ``
		}
		users = append(users, u)
	}
	result.Users = users
	return result, nil
}
