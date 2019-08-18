package nps

// Org ...
type Org struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// OrgsResult ...
type OrgsResult struct {
	TotalCount int    `json:"totalCount"`
	Orgs       []*Org `json:"orgs"`
}
