package main

import (
	"WinterSoldier/config/env"

	mongoDB "gitlab.com/q2tech/q2-stormbreaker/database/mongodb"
)

// Collections
var (
	peopleCollection mongoDB.CollectionName = "people"
)

func main() {
	env.GetEnvironment()

	collections := []mongoDB.CollectionName{peopleCollection}

	mongo := mongoDB.New(env.MongoEnv.URI, env.MongoEnv.DB, collections)

	mongo.Start()
}
