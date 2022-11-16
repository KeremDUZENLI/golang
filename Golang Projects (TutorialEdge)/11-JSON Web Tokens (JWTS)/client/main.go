package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"

	jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte(os.Getenv("SECRET"))

func homePage(w http.ResponseWriter, r *http.Request) {

	validToken, _ := GenerateJWT()
	client := &http.Client{}

	req, _ := http.NewRequest("GET", "http://localhost:9002/", nil)
	req.Header.Set("Token", validToken)

	res, _ := client.Do(req)
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Fprintf(w, string(body))
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

	tokenString, _ := token.SignedString(mySigningKey)

	return tokenString, nil
}

func main() {
	godotenv.Load(".env")

	fmt.Println("My Simple Client")

	handleRequests()
}
