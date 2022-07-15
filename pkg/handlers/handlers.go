package handlers

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"main.go/pkg/user"
)

var ErrorMethodNotAllowed = "Method not allowed"

type ErrorBody struct {
	ErrorMsg *string `json:"error, omitempty"`
}

func GetUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (events.APIGatewayProxyResponse, error) {
	email := req.QueryStringParamters["email"]
	if len(email) == 0 {
		result, err := user.FetchUser(email, tableName, dynaClient)
		if err != nil {
			return apiResponse(http.StatusMethodNotAllowed, ErrorBody{aws.String(err.Error())})
		}
		return apiResponse(http.StatusOK, result)

	}

	result, err := user.FetchUser(tableName, dynaClient)
	if err != nil {
		return apiResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}
	return apiResponse(http.StatusOK, result)
	
}

func CreateUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (events.APIGatewayProxyResponse, error) {

}

func UpdateUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (events.APIGatewayProxyResponse, error) {

}

func DeleteUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (events.APIGatewayProxyResponse, error) {

}

func Unhandle() (*events.APIGatewayProxyResponse, error) { 
	return apiResponse(http.StatusMethodNotAllowed, ErrorMethodNotAllowed)
}