package nps

// User ...
type User struct {
	ID         int         `json:"id"`
	Email      *string     `json:"email"`
	FirstName  *string     `json:"firstName" db:"first_name"`
	LastName   *string     `json:"lastName" db:"last_name"`
	NickName   *string     `json:"nickName" db:"nick_name"`
	Orgs       []*Org      `json:"orgs"`
	Tags       []*Tag      `json:"tags"`
	Categories []*Category `json:"categories"`
}
