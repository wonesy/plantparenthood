package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/wonesy/plantparenthood/graph/generated"
	"github.com/wonesy/plantparenthood/graph/model"
	"github.com/wonesy/plantparenthood/internal/pkg/db"
	"github.com/wonesy/plantparenthood/pkg/jwt"
)

func (r *mutationResolver) CreateMember(ctx context.Context, input model.NewMember) (*model.Member, error) {
	return r.memberHandler.Create(&input)
}

func (r *mutationResolver) CreatePlant(ctx context.Context, input model.NewPlant) (*model.Plant, error) {
	_, err := r.memberHandler.ValidateMemberFromContext(ctx)
	if err != nil {
		return nil, err
	}

	return r.plantHandler.Create(&input)
}

func (r *mutationResolver) CreateCareRegimen(ctx context.Context, input model.NewCareRegimen) (*model.CareRegimen, error) {
	_, err := r.memberHandler.ValidateMemberFromContext(ctx)
	if err != nil {
		return nil, err
	}

	return r.careRegimenHandler.Create(&input)
}

func (r *mutationResolver) AddPlantToNursery(ctx context.Context, input model.NewNurseryAddition) (*model.PlantBaby, error) {
	memberID, err := r.memberHandler.ValidateMemberFromContext(ctx)
	if err != nil {
		return nil, err
	}

	// add plant to member's nursery
	pb, err := r.plantBabyHandler.Create(memberID, &input)
	if err != nil {
		return nil, err
	}

	pb.Plant, err = r.plantHandler.GetByID(input.PlantID)
	if err != nil {
		return nil, err
	}

	return pb, nil
}

func (r *mutationResolver) CreateWatering(ctx context.Context, input *model.NewWatering) (*model.Watering, error) {
	_, err := r.memberHandler.ValidateMemberFromContext(ctx)
	if err != nil {
		return nil, err
	}

	return r.wateringHandler.Create(input)
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

func (r *queryResolver) MemberByID(ctx context.Context, id string) (*model.Member, error) {
	return r.memberHandler.GetByID(id)
}

func (r *queryResolver) Plants(ctx context.Context) ([]*model.Plant, error) {
	return r.plantHandler.GetAll()
}

func (r *queryResolver) PlantByID(ctx context.Context, id string) (*model.Plant, error) {
	return r.plantHandler.GetByID(id)
}

func (r *queryResolver) Nursery(ctx context.Context) ([]*model.PlantBaby, error) {
	memberID, err := r.memberHandler.ValidateMemberFromContext(ctx)
	if err != nil {
		return nil, err
	}

	return r.plantBabyHandler.GetAllByOwnerID(memberID)
}

func (r *queryResolver) PlantFromNursery(ctx context.Context, id string) (*model.PlantBaby, error) {
	return r.plantBabyHandler.GetByID(id)
}

func (r *queryResolver) CareRegimens(ctx context.Context) ([]*model.CareRegimen, error) {
	return r.careRegimenHandler.GetAll()
}

func (r *queryResolver) CareRegimenByID(ctx context.Context, id string) (*model.CareRegimen, error) {
	return r.careRegimenHandler.GetByID(id)
}

func (r *queryResolver) Waterings(ctx context.Context, plantBabyID string) ([]*model.Watering, error) {
	memberID, err := r.memberHandler.ValidateMemberFromContext(ctx)
	if err != nil {
		return nil, err
	}

	// ensure that the member from the auth token owns this plant
	if err := r.plantBabyHandler.MemberOwnsPlant(memberID, plantBabyID); err != nil {
		if _, ok := err.(*db.NoSuchEntity); ok {
			return nil, errors.New("Member does not own requested plant")
		}
		return nil, err
	}

	return r.wateringHandler.GetByPlantBabyID(plantBabyID)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
