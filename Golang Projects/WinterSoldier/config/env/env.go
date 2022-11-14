package env

import (
	"gitlab.com/q2tech/q2-stormbreaker/env"
)

type AppEnvironment struct {
	Port       string `env:"SERVER_PORT"`
	CorsOrigin string `env:"CORS_ORIGIN"`
}

type MongoEnvironment struct {
	URI string `env:"MONGO_URI"`
	DB  string `env:"MONGO_DB"`
}

var AppEnv AppEnvironment
var MongoEnv MongoEnvironment

func GetEnvironment() {
	env.LoadEnv(&AppEnv, &MongoEnv)
}
