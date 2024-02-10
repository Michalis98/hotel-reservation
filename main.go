package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/Michalis98/hotel-reservation/api"
	"github.com/Michalis98/hotel-reservation/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dburi = "mongodb://localhost:27017"
const dbname = "hotel-reservation"
const userColl = "users"

func main() {
	//Database stuff
	//Connect to the database
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}
	//Crate a new context
	ctx := context.Background()
	//Create a new database and a new collection
	coll := client.Database(dbname).Collection(userColl)
	//Crate a new user
	user := types.User{
		FirstName: "James",
		LastName:  "Ioannou",
	}
	//Insert the user to the database
	_, err = coll.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}
	//Find the user ( the query is empty because it will return the 1st occurence in the database, in our case is James)
	var james types.User
	if err := coll.FindOne(ctx, bson.M{}).Decode(&james); err != nil {
		log.Fatal(err)
	}
	fmt.Print(james)

	listenAddr := flag.String("listenAddr", ":5000", "The listen address of the API server")
	flag.Parse()
	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	apiv1.Get("/user", api.HandleGetUsers)
	apiv1.Get("/user/:id", api.HandleGetUser)

	app.Listen(*listenAddr)

}
