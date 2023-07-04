package main

import (
	"context"
	"fmt"
	"log"

	"github.com/diazharizky/go-mongodb-with-tests/internal/app"
	"github.com/diazharizky/go-mongodb-with-tests/internal/models"
	"github.com/diazharizky/go-mongodb-with-tests/internal/repositories"
	"github.com/diazharizky/go-mongodb-with-tests/pkg/mongodb"
)

var appCtx *app.Context

func init() {
	initAppCtx()
}

func main() {
	defer appCtx.DBClient.Disconnect(context.TODO())

	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	newUser := models.User{
		Username: "foo",
		Email:    "foo@mail.com",
		Age:      20,
		FullName: "Foo To Risk Tech",
	}

	id, err := appCtx.UserRepository.Create(ctx, newUser)
	if err != nil {
		fmt.Printf("Error unable to create new user: %v\n", err)
	}

	users, err := appCtx.UserRepository.List(ctx)
	if err != nil {
		fmt.Printf("Error unable to list users: %v\n", err)
	}

	fmt.Println("User list")
	for _, u := range users {
		fmt.Println("-- User")
		fmt.Println("---- Username", u.Username)
		fmt.Println("---- Email", u.Email)
		fmt.Println("---- Age", u.Age)
		fmt.Println("---- Full Name", u.FullName)
	}

	err = appCtx.UserRepository.Delete(ctx, *id)
	if err != nil {
		fmt.Printf("Error unable to deleted existing user: %v\n", err)
	}
}

func initAppCtx() {
	var err error

	appCtx = &app.Context{}

	appCtx.DBClient, err = mongodb.GetClient()
	if err != nil {
		log.Fatalf("Error unable to get DB client: %v", err)
	}

	db := appCtx.DBClient.Database("go_mongodb")
	appCtx.UserRepository = repositories.NewUserRepository(db)
}
