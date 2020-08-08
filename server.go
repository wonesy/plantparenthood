package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"

	"github.com/wonesy/plantparenthood/graph"
	"github.com/wonesy/plantparenthood/graph/generated"
	"github.com/wonesy/plantparenthood/internal/auth"
	ppcareregimen "github.com/wonesy/plantparenthood/internal/careregimen"
	ppmember "github.com/wonesy/plantparenthood/internal/member"
	database "github.com/wonesy/plantparenthood/internal/pkg/db"
	ppplant "github.com/wonesy/plantparenthood/internal/plant"
	ppplantbaby "github.com/wonesy/plantparenthood/internal/plantbaby"
	ppwatering "github.com/wonesy/plantparenthood/internal/watering"
)

const defaultPort = "8080"

func main() {
	user := "pp"
	password := "password"
	dbname := "plantparenthood"

	conn, err := database.Open(user, password, dbname)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	port := os.Getenv("PP_PORT")
	if port == "" {
		port = defaultPort
	}

	// initialize all model handlers
	memberHandler := ppmember.NewHandler(conn)
	plantHandler := ppplant.NewHandler(conn)
	plantBabyHandler := ppplantbaby.NewHandler(conn)
	careRegimenHandler := ppcareregimen.NewHandler(conn)
	wateringHandler := ppwatering.NewHandler(conn)

	/************************************************/
	/************************************************/
	/************************************************/
	/************************************************/
	/************************************************/
	/************************************************/
	// seed database, this is for testing and early development
	seed(memberHandler, careRegimenHandler, plantHandler, plantBabyHandler, wateringHandler)
	/************************************************/
	/************************************************/
	/************************************************/
	/************************************************/
	/************************************************/
	/************************************************/

	// initialize the resolver
	resolver := graph.NewResolver(conn, memberHandler, plantHandler, careRegimenHandler, plantBabyHandler, wateringHandler)

	// create gql server
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	router := chi.NewRouter()
	router.Use(auth.CheckTokenMiddleware())

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
