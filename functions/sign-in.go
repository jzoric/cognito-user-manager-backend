package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
)

type Response struct {
	Token string `json:"token"`
}

type SignIn struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func credentialsValid(username, password string) bool {
	return username == os.Getenv("USERNAME") && password == os.Getenv("PASSWORD")
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var si SignIn
	json.Unmarshal([]byte(request.Body), &si)

	if !credentialsValid(si.Username, si.Password) {
		return events.APIGatewayProxyResponse{Body: "Wrong credentials", StatusCode: http.StatusForbidden}, nil
	}

	token := jwt.New(jwt.SigningMethodHS256)

	tokenString, err := token.SignedString([]byte("secret2018!"))
	if err != nil {
		return events.APIGatewayProxyResponse{Body: "Can't create token", StatusCode: http.StatusInternalServerError}, nil
	}
	t, _ := json.Marshal(&Response{tokenString})
	return events.APIGatewayProxyResponse{Body: string(t), Headers: map[string]string{"Access-Control-Allow-Origin": "*"}, StatusCode: http.StatusOK}, nil
}

func main() {
	lambda.Start(Handler)
}
