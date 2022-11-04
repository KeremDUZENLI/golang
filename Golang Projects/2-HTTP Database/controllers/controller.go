package controller

import (
	"encoding/json"
	"fmt"
	model "http/models"
	util "http/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var NewUser model.User

func GetURL(writer http.ResponseWriter, reader *http.Request) {
	fmt.Fprintf(writer, "GET USER ALL\n")

	userAll := model.GetUser()
	writings, _ := json.Marshal(userAll)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(writings)
}

func GetIDURL(writer http.ResponseWriter, reader *http.Request) {
	fmt.Fprintf(writer, "GET USER ID\n")

	url := mux.Vars(reader)
	urlID := url["ID"]
	numID, _ := strconv.ParseInt(urlID, 0, 0)

	userID, _ := model.GetUserID(numID)
	writings, _ := json.Marshal(userID)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(writings)
}

func PostURL(writer http.ResponseWriter, reader *http.Request) {
	fmt.Fprintf(writer, "POST USER\n")

	userCreate := &model.User{}
	util.ParseBody(reader, userCreate)

	x := userCreate.PostUser()
	writings, _ := json.Marshal(x)

	writer.WriteHeader(http.StatusOK)
	writer.Write(writings)
}

func DeleteURL(writer http.ResponseWriter, reader *http.Request) {
	fmt.Fprintf(writer, "DELETE USER\n")

	url := mux.Vars(reader)
	urlID := url["ID"]
	numID, _ := strconv.ParseInt(urlID, 0, 0)

	// Returns which is gonna be deleted
	userIDDeleted, _ := model.GetUserID(numID)
	writingsDeleted, _ := json.Marshal(userIDDeleted)

	// Deletes
	userID := model.DeleteUser(numID)
	json.Marshal(userID)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(writingsDeleted)
}
