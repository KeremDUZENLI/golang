package main

import (
	"postgre-project/common/env"
	"postgre-project/database"
)

func main() {
	env.Load()

	database.ConnectDB()
	database.LoadDB()
	database.CloseDB()
}
