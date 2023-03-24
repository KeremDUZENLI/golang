package main

import (
	"fmt"
	"postgre-project/common/env"
	"postgre-project/database"
	"postgre-project/database/model"
	"postgre-project/dto"
	"postgre-project/repository"
)

func main() {
	env.Load()

	database.ConnectDB()

	LoadDB()

	repository.AddUser(dto.DtoSignUp{
		Password:     "bbb",
		Token:        "",
		RefreshToken: "",
		FirstName:    "",
		LastName:     "",
		Email:        "aaa",
		UserType:     "",
	})

	res, err := repository.FindUserByEmailPassword(dto.DtoLogIn{
		Email:    "xxx",
		Password: "",
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("\nEmail: %s \nPassword: %s\n\n", res.Email, res.Password)

	database.CloseDB()
}

func LoadDB() {
	database.Instance.AutoMigrate(&model.Tables{})
}
