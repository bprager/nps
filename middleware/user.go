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
	SELECT 
( SELECT count(*) FROM users ) AS totalCount,
users.id AS userID, users.email, users.first_name AS firstName, users.last_name AS lastName, users.nick_name AS nickName,
tags.id AS tagID, tags.name AS tagName, tags."attribute", tags."number", tags."timestamp",
orgs."id" AS orgID, orgs."name" AS orgName,
categories."id" AS catID, categories."name" AS category_name, categories.parent
FROM users 
LEFT OUTER JOIN user_tag_idx AS tidx ON tidx."user" = users.id
LEFT OUTER JOIN tags ON tidx.tag = tags.id
LEFT OUTER JOIN user_org_idx AS oidx ON oidx."user" = users.id
LEFT OUTER JOIN orgs ON oidx.org = orgs.id
LEFT OUTER JOIN user_category_idx AS cidx ON cidx."user" = users.id
LEFT OUTER JOIN categories ON cidx.category = categories.id
WHERE users.id IN (SELECT users.id FROM users LIMIT ? OFFSET ?)
`
	rows, err := DB.Queryx(query, limit, offset)
	if err != nil {
		log.Fatalln(err)
	}
	result := new(UsersResult)
	var users []*User
	// var orgs []*Org
	// var tags []*Tag
	// var cats []*Category

	for rows.Next() {
		u := new(User)
		o := new(NullInt64)
		t := new(NullInt64)
		c := new(NullInt64)

		var orgs []Org
		var tags []Tag
		var cats []Category

		err := rows.Scan(&u.ID, &u.Email, &u.FirstName, &u.LastName, &u.NickName, o, t, c, &result.TotalCount)
		if err != nil {
			log.Fatalln(err)
		}
		// TODO: optimize looping
		// Get tags, orgs and categories if there are any
		if o.Valid || t.Valid || c.Valid {
			query := `
			SELECT tags."id" AS tagID, tags."name" AS tagName, tags."attribute", tags."number", tags."timestamp",
			orgs."id" AS orgID, orgs."name" AS org_name, 
			categories."id" as catID, categories."name" AS catName
			FROM users
			LEFT OUTER JOIN user_tag_idx AS tidx ON tidx."user" = users.id
			LEFT OUTER JOIN tags ON tidx.tag = tags.id
			LEFT OUTER JOIN user_org_idx AS oidx ON oidx."user" = users.id
			LEFT OUTER JOIN orgs ON oidx.org = orgs.id
			LEFT OUTER JOIN user_category_idx AS cidx on cidx."user" = users.id
			LEFT OUTER JOIN categories ON cidx.category = categories.id
			WHERE users.id = ?`

			rows, err := DB.Queryx(query, u.ID)
			if err != nil {
				log.Fatalln(err)
			}

			for rows.Next() {
				tagID := new(NullInt64)
				tagName := new(NullString)
				attribute := new(NullString)
				number := new(NullInt64)
				timestamp := new(NullString)
				orgID := new(NullInt64)
				orgName := new(NullString)
				catID := new(NullInt64)
				catName := new(NullString)
				err := rows.Scan(&tagID, &tagName, &attribute, &number, &timestamp,
					&orgID, &orgName, &catID, &catName)
				if err != nil {
					log.Fatalln(err)
				}
				// check if a tag is present
				if tagID.Valid {
					t := new(Tag)
					if tagName.Valid {
						t.ID = int(tagID.Int64)
						// check if tag already exist for user
						new := true
						for _, n := range tags {
							if n.ID == t.ID {
								new = false
							}
						}
						if new {
							if tagName.Valid {
								t.Name = tagName.String
							}
							t.Attribute = attribute
							t.Number = number
							t.Timestamp = timestamp
							tags = append(tags, *t)
						}
					}
				}
				// check if an org is present
				if orgID.Valid {
					o := new(Org)
					if orgName.Valid {
						o.ID = int(orgID.Int64)
						new := true
						for _, n := range tags {
							if n.ID == o.ID {
								new = false
							}
						}
						if new {
							if orgName.Valid {
								o.Name = orgName.String
							}
							orgs = append(orgs, *o)
						}
					}
				}
			}

		}
		u.Orgs = &orgs
		u.Tags = &tags
		u.Categories = &cats
		users = append(users, u)
	}
	result.Users = users
	return result, nil
}
