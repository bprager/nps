// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package nps

type CategoriesResult struct {
	TotalCount int         `json:"totalCount"`
	Categories []*Category `json:"categories"`
}

type Category struct {
	ID     string    `json:"id"`
	Name   string    `json:"name"`
	Parent *Category `json:"parent"`
}

type Note struct {
	ID    string `json:"id"`
	Text  string `json:"text"`
	Saved string `json:"saved"`
}

type Question struct {
	ID   string `json:"id"`
	Open bool   `json:"open"`
	Body string `json:"body"`
}

type Survey struct {
	ID            string    `json:"id"`
	Start         *string   `json:"start"`
	End           *string   `json:"end"`
	ScoreQuestion *Question `json:"scoreQuestion"`
	OpenQuestion  *Question `json:"openQuestion"`
	Note          *Note     `json:"note"`
}

type SurveysResult struct {
	TotalCount int       `json:"totalCount"`
	Surveys    []*Survey `json:"surveys"`
}

type TagsResult struct {
	TotalCount int    `json:"totalCount"`
	Tags       []*Tag `json:"tags"`
}

type UsersResult struct {
	TotalCount int     `json:"totalCount"`
	Users      []*User `json:"users"`
}
