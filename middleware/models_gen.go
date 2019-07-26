// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package main

type Category struct {
	ID     string    `json:"id"`
	Name   *string   `json:"name"`
	Parent *Category `json:"parent"`
}

type Note struct {
	ID    string  `json:"id"`
	Text  *string `json:"text"`
	Saved string  `json:"saved"`
}

type Org struct {
	ID   string  `json:"id"`
	Name *string `json:"name"`
}

type Question struct {
	ID   string  `json:"id"`
	Open *bool   `json:"open"`
	Body *string `json:"body"`
}

type Survey struct {
	ID            string    `json:"id"`
	Start         *string   `json:"start"`
	End           *string   `json:"end"`
	ScoreQuestion *Question `json:"scoreQuestion"`
	OpenQuestion  *Question `json:"openQuestion"`
	Note          *Note     `json:"note"`
}

type Tag struct {
	ID        string  `json:"id"`
	Name      *string `json:"name"`
	Attribute *string `json:"attribute"`
	Number    *int    `json:"number"`
	Timestamp *string `json:"timestamp"`
}

type User struct {
	ID         string      `json:"id"`
	Email      *string     `json:"email"`
	FirstName  *string     `json:"firstName"`
	LastName   *string     `json:"lastName"`
	NickName   *string     `json:"nickName"`
	Orgs       []*Org      `json:"orgs"`
	Tags       []*Tag      `json:"tags"`
	Categories []*Category `json:"categories"`
}