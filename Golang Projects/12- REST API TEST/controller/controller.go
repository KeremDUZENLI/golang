package controller

import (
	"net/http"
	"restApiTest/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Begin(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "CONNECTED")
}

func Create(c *gin.Context) {
	var newUser model.User
	c.BindJSON(&newUser)
	model.Users = append(model.Users, newUser)

	c.IndentedJSON(http.StatusCreated, newUser)
}

func Read(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, model.Users)
}

func Update(c *gin.Context) {
	nStr := c.Param("employeeID")
	n, _ := strconv.Atoi(nStr)

	var updatedUser model.User
	c.BindJSON(&updatedUser)

	for index, user := range model.Users {
		if user.EmployeeID == n {
			model.Users[index].EmployeeID = updatedUser.EmployeeID
			model.Users[index].Name = updatedUser.Name
			model.Users[index].Email = updatedUser.Email
			model.Users[index].Document = updatedUser.Document
			model.Users[index].RegisteredAt = updatedUser.RegisteredAt
			break
		}
	}

	// c.IndentedJSON(http.StatusOK, updatedUser)
	c.IndentedJSON(http.StatusOK, model.Users)
}

func Delete(c *gin.Context) {
	nStr := c.Param("employeeID")
	n, _ := strconv.Atoi(nStr)

	for index, user := range model.Users {
		if user.EmployeeID == n {
			c.IndentedJSON(http.StatusOK, user)
			model.Users = append(model.Users[:index], model.Users[index+1:]...)
			break
		}
	}
}

// DB ----------------------------------------------------------------

func BeginDatabase(c *gin.Context) {
	userBegin := &model.User{}
	c.ShouldBindJSON(userBegin)

	c.IndentedJSON(http.StatusOK, userBegin.DatabaseBeginUser())
}

func CreateDatabase(c *gin.Context) {
	userCreate := &model.User{}
	c.ShouldBindJSON(userCreate)

	c.IndentedJSON(http.StatusOK, userCreate.DatabaseCreateUser())
}

func ReadDatabase(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, model.DatabaseReadUser())
}

func UpdateDatabase(c *gin.Context) {
	model.DatabaseUpdateUser()
}

func DeleteDatabase(c *gin.Context) {
	nStr := c.Param("employeeID")
	n, _ := strconv.Atoi(nStr)

	for _, user := range *model.DatabaseReadUser() {
		if user.EmployeeID == n {
			c.IndentedJSON(http.StatusOK, user)
			model.DatabaseDeleteUser(n)
			break
		}
	}
}
