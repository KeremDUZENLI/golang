package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"

	jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte(os.Getenv("SECRET"))

func homePage(w http.ResponseWriter, r *http.Request) {

	validToken, err := GenerateJWT()
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	fmt.Fprintf(w, validToken)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":9001", nil))
}

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "Orkut PERUU"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("ERROR: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func main() {
	godotenv.Load(".env")

	fmt.Println("My Simple Client")

	// tokenString, err := GenerateJWT()
	// if err != nil {
	// 	fmt.Println("Error!")
	// }
	// fmt.Println(tokenString)

	handleRequests()
}
