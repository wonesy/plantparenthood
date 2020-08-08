package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/wonesy/plantparenthood/graph/generated"
	"github.com/wonesy/plantparenthood/graph/model"
	"github.com/wonesy/plantparenthood/internal/auth"
	"github.com/wonesy/plantparenthood/pkg/jwt"
)

func (r *mutationResolver) CreateMember(ctx context.Context, input model.NewMember) (string, error) {
	id, err := r.memberHandler.Create(&input)
	if err != nil {
		return "", err
	}

	return jwt.GenerateToken(id)
}

func (r *mutationResolver) CreatePlant(ctx context.Context, input model.NewPlant) (string, error) {
	id := auth.IDFromContext(ctx)
	if id == "" {
		return "", &auth.UnauthenticatedError{}
	}

	return r.plantHandler.Create(&input)
}

func (r *mutationResolver) AddPlantToNursery(ctx context.Context, input model.NewNurseryAddition) (*model.PlantBaby, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	id, err := r.memberHandler.Login(&input)
	if err != nil {
		return "", err
	}

	return jwt.GenerateToken(id)
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshToken) (string, error) {
	id, err := jwt.ParseToken(input.Token)
	if err != nil {
		return "", err
	}

	return jwt.GenerateToken(id)
}

func (r *queryResolver) Members(ctx context.Context) ([]*model.Member, error) {
	return r.memberHandler.GetAll()
}

func (r *queryResolver) Plants(ctx context.Context) ([]*model.Plant, error) {
	return r.plantHandler.GetAll()
}

func (r *queryResolver) Nursery(ctx context.Context) ([]*model.PlantBaby, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
