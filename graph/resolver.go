package graph

import (
	"database/sql"

	ppcareregimen "github.com/wonesy/plantparenthood/internal/careregimen"
	ppmember "github.com/wonesy/plantparenthood/internal/member"
	ppplant "github.com/wonesy/plantparenthood/internal/plant"
	ppplantbaby "github.com/wonesy/plantparenthood/internal/plantbaby"
	ppwatering "github.com/wonesy/plantparenthood/internal/watering"
)

// Resolver graphql type resolver
type Resolver struct {
	db                 *sql.DB
	memberHandler      *ppmember.Handler
	plantHandler       *ppplant.Handler
	careRegimenHandler *ppcareregimen.Handler
	plantBabyHandler   *ppplantbaby.Handler
	wateringHandler    *ppwatering.Handler
}

// NewResolver constructor for Resolver
func NewResolver(
	db *sql.DB,
	mh *ppmember.Handler,
	ph *ppplant.Handler,
	crh *ppcareregimen.Handler,
	pbh *ppplantbaby.Handler,
	wh *ppwatering.Handler,
) *Resolver {
	return &Resolver{
		db:                 db,
		memberHandler:      mh,
		plantHandler:       ph,
		careRegimenHandler: crh,
		plantBabyHandler:   pbh,
		wateringHandler:    wh,
	}
}
