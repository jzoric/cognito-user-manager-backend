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

func Handler(requst events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	sess := session.Must(session.NewSession(&aws.Config{Region: aws.String("eu-central-1")}))
	cip := cognitoidentityprovider.New(sess)
	result, err := cip.ListUserPools(&cognitoidentityprovider.ListUserPoolsInput{MaxResults: aws.Int64(60)})
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: http.StatusInternalServerError}, nil
	}
	userPools, _ := json.Marshal(result.UserPools)
	return events.APIGatewayProxyResponse{Body: string(userPools), Headers: map[string]string{"Access-Control-Allow-Origin": "*"}, StatusCode: http.StatusOK}, nil
}

func main() {
	lambda.Start(Handler)
}
