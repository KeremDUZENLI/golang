package model

import "restApiTest/configs"

type User struct {
	EmployeeID   int    `json:"employeeID"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Document     string `json:"document"`
	RegisteredAt string `json:"registeredAt"`
}

var Users = []User{
	{EmployeeID: 1, Name: "A", Email: "AA@", Document: "AAA", RegisteredAt: ""},
	{EmployeeID: 2, Name: "B", Email: "BB@", Document: "BBB", RegisteredAt: ""},
	{EmployeeID: 3, Name: "C", Email: "CC@", Document: "CCC", RegisteredAt: ""},
}

// DB ----------------------------------------------------------------

func (aStruct *User) DatabaseBeginUser() *[]User {
	for _, user := range Users {
		if configs.Database.Where("employee_id = ?", user.EmployeeID).First(&user).RecordNotFound() {
			configs.Database.Create(&user)
		}
	}

	return &Users
}

func (aStruct *User) DatabaseCreateUser() *User {
	if configs.Database.Where("employee_id = ?", aStruct.EmployeeID).First(&aStruct).RecordNotFound() {
		configs.Database.Create(&aStruct)
	}

	return aStruct
}

func DatabaseReadUser() *[]User {
	var usersAll []User
	configs.Database.Find(&usersAll)
	return &usersAll
}

func DatabaseUpdateUser() {

}

func DatabaseDeleteUser(id int) User {
	var userDelete User
	configs.Database.Where("employee_id = ?", id).Delete(userDelete)
	return userDelete
}
