package nps

import (
	"context"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Surveys(ctx context.Context) ([]*Survey, error) {
	panic("not implemented")
}
func (r *queryResolver) Users(ctx context.Context) ([]*User, error) {
	panic("not implemented")
}
func (r *queryResolver) Orgs(ctx context.Context) ([]*Org, error) {
	panic("not implemented")
}
func (r *queryResolver) Categories(ctx context.Context) ([]*Category, error) {
	panic("not implemented")
}
func (r *queryResolver) Tags(ctx context.Context) ([]*Tag, error) {
	panic("not implemented")
}
