package main

import (
	"github.com/wonesy/plantparenthood/graph/model"
	ppcareregimen "github.com/wonesy/plantparenthood/internal/careregimen"
	ppmember "github.com/wonesy/plantparenthood/internal/member"
	ppplant "github.com/wonesy/plantparenthood/internal/plant"
	ppplantbaby "github.com/wonesy/plantparenthood/internal/plantbaby"
	ppwatering "github.com/wonesy/plantparenthood/internal/watering"
)

func seed(m *ppmember.Handler, c *ppcareregimen.Handler, p *ppplant.Handler, pb *ppplantbaby.Handler, w *ppwatering.Handler) {
	// create two users
	members, _ := m.GetAll()
	if len(members) == 0 {
		memberA, _ := m.Create(&model.NewMember{
			Email:    "a@a.a",
			Password: "a",
		})

		memberB, _ := m.Create(&model.NewMember{
			Email:    "b@b.b",
			Password: "b",
		})
		members = append(members, []*model.Member{memberA, memberB}...)
	}

	plants, _ := p.GetAll()
	if len(plants) == 0 {
		// create two plants
		pltA, _ := p.Create(&model.NewPlant{
			CommonName:      "common plant a",
			BotanicalName:   "botanical plant a",
			SoilPreference:  "soil a",
			WaterPreference: "water a",
			SunPreference:   "sun a",
		})

		pltB, _ := p.Create(&model.NewPlant{
			CommonName:      "common plant b",
			BotanicalName:   "botanical plant b",
			SoilPreference:  "soil b",
			WaterPreference: "water b",
			SunPreference:   "sun b",
		})

		plants = append(plants, []*model.Plant{pltA, pltB}...)
	}

	// care regimen

	cr, _ := c.Create(&model.NewCareRegimen{
		Waterhr: 150,
		Waterml: 200,
	})

	// own plants
	babies, _ := pb.GetAllByOwnerID(members[0].ID)
	babies2, _ := pb.GetAllByOwnerID(members[1].ID)
	babies = append(babies, babies2...)
	if len(babies) == 0 {
		pb1, _ := pb.Create(members[0].ID, &model.NewNurseryAddition{
			CareRegimenID: cr.ID,
			Nickname:      "orange",
			Location:      "living room",
			PlantID:       plants[0].ID,
		})

		pb2, _ := pb.Create(members[1].ID, &model.NewNurseryAddition{
			CareRegimenID: cr.ID,
			Nickname:      "blueberry",
			Location:      "living room",
			PlantID:       plants[0].ID,
		})

		pb3, _ := pb.Create(members[1].ID, &model.NewNurseryAddition{
			CareRegimenID: cr.ID,
			Nickname:      "blackberry",
			Location:      "living room",
			PlantID:       plants[1].ID,
		})

		babies = append(babies, []*model.PlantBaby{pb1, pb2, pb3}...)
	}

	// waterings
	_, _ = w.Create(&model.NewWatering{
		PlantBabyID: babies[0].ID,
		Amountml:    100,
		WateredOn:   "2019-11-22T07:20:50.52Z",
	})

	_, _ = w.Create(&model.NewWatering{
		PlantBabyID: babies[0].ID,
		Amountml:    200,
	})
}
