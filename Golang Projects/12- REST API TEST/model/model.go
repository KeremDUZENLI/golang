package model

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
