package repository

import (
	"errors"
	"postgre-project/database"
	"postgre-project/database/model"
	"postgre-project/dto"
	"postgre-project/dto/mapper"
	"postgre-project/middleware"
	"postgre-project/middleware/token"
)

// SignUp
func AddUser(dSU dto.DtoSignUp) {
	aMap := mapper.MapperSignUp(&dSU)

	setValues(&aMap)

	database.Instance.Create(&aMap)
}

func setValues(person *model.Tables) error {
	person.Password, _ = middleware.HashPassword(person.Password)
	_, errPassword := middleware.HashPassword(person.Password)

	token, refreshToken, errToken := token.GenerateToken(person.FirstName, person.LastName, person.Email, person.UserType)
	person.Token = token
	person.RefreshToken = refreshToken

	if errPassword != nil && errToken != nil {
		return errors.New("error setValues")
	}

	return nil
}

// LogIn
func FindUserByEmailPassword(dLI dto.DtoLogIn) (model.Tables, error) {
	aMap := mapper.MapperLogIn(&dLI)

	err := database.Instance.Where("email = ?", aMap.Email).
		Where("password = ?", aMap.Password).Error

	if err != nil {
		return model.Tables{}, err
	}

	return aMap, nil
}
