package nps

import (
	"context"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

// Resolver ...
type Resolver struct{}

// Mutation ...
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

// Query ...
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) AddOrg(ctx context.Context, name string) (bool, error) {
	panic("not implemented")
}
func (r *mutationResolver) AddCategory(ctx context.Context, name string, parent *string) (bool, error) {
	panic("not implemented")
}
func (r *mutationResolver) AddTag(ctx context.Context, name string, attribute *string, number *int, timestamp *string) (bool, error) {
	panic("not implemented")
}
func (r *mutationResolver) AddUser(ctx context.Context, email *string, firstName *string, lastName *string, nickName *string) (bool, error) {
	panic("not implemented")
}
func (r *mutationResolver) ChangeUser(ctx context.Context, email *string, firstName *string, lastName *string, nickName *string) (bool, error) {
	panic("not implemented")
}
func (r *mutationResolver) ChangeCategory(ctx context.Context, name string, parent *string) (bool, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Survey(ctx context.Context, id string) (*Survey, error) {
	panic("not implemented")
}
func (r *queryResolver) AllSurveys(ctx context.Context) ([]*Survey, error) {
	panic("not implemented")
}
func (r *queryResolver) User(ctx context.Context, id string) (*User, error) {
	panic("not implemented")
}
func (r *queryResolver) Users(ctx context.Context, tags []string, categories []string, org string) ([]*User, error) {
	panic("not implemented")
}
func (r *queryResolver) AllOrgs(ctx context.Context) ([]*Org, error) {
	panic("not implemented")
}
func (r *queryResolver) AllCategories(ctx context.Context) ([]*Category, error) {
	panic("not implemented")
}
func (r *queryResolver) Tag(ctx context.Context, id string) (*Tag, error) {
	panic("not implemented")
}
func (r *queryResolver) Tags(ctx context.Context, user string) ([]*Tag, error) {
	panic("not implemented")
}
