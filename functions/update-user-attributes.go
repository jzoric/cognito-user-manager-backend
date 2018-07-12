package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"net/http"
)

type Request struct {
	UserPoolId     string                                   `json:"user_pool_id"`
	Username       string                                   `json:"username"`
	UserAttributes []*cognitoidentityprovider.AttributeType `json:"user_attributes"`
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var r Request
	json.Unmarshal([]byte(request.Body), &r)

	sess := session.Must(session.NewSession(&aws.Config{Region: aws.String("eu-central-1")}))

	cip := cognitoidentityprovider.New(sess)
	_, err := cip.AdminUpdateUserAttributes(&cognitoidentityprovider.AdminUpdateUserAttributesInput{UserPoolId: aws.String(r.UserPoolId), Username: aws.String(r.Username), UserAttributes: r.UserAttributes})
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: http.StatusInternalServerError}, nil
	}

	return events.APIGatewayProxyResponse{Headers: map[string]string{"Access-Control-Allow-Origin": "*"}, StatusCode: http.StatusOK}, nil
}

func main() {
	lambda.Start(Handler)
}
