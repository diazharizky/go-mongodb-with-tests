package app

import "go.mongodb.org/mongo-driver/mongo"

type Context struct {
	DBClient *mongo.Client

	UserRepository IUserRepository
}
