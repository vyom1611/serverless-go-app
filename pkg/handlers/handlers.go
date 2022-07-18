package handlers

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"main.go/pkg/user"
)

//Defining customer error
var ErrorMethodNotAllowed = "method not allowed"

//Custom error body structure
type ErrorBody struct{
	ErrorMsg *string `json:"error,omitempty"`
}


//GET user function
func GetUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI)(
	*events.APIGatewayProxyResponse, error,
){

	//Finding email from request body and validating it
	email := req.QueryStringParameters["email"]
	if len(email) > 0 {
		result, err := user.FetchUser(email, tableName, dynaClient)
		if err!= nil {
			return apiResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
		}
		return apiResponse(http.StatusOK, result)
	}

	//Extracting users from dynaclient and saving to result (using fetch method defined in user.go)
	result, err := user.FetchUsers(tableName, dynaClient)
	if err!= nil {
		return apiResponse(http.StatusBadRequest, ErrorBody{
			aws.String(err.Error()),
		})
	}
	return apiResponse(http.StatusOK, result)

}


//CREATE method
func CreateUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI)(
	*events.APIGatewayProxyResponse, error,
){
	//Using create method from user.go and passing in the current request/tablename and dyna client
	result, err := user.CreateUser(req, tableName, dynaClient)
	if err!=nil {
		return apiResponse(http.StatusBadRequest, ErrorBody{
			aws.String(err.Error()),
		})
	}
	return apiResponse(http.StatusCreated, result)
}


//UPDATE method
func UpdateUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI)(
	*events.APIGatewayProxyResponse, error,
){
	//Using update method in user.go to current params
	result, err := user.UpdateUser(req, tableName, dynaClient)
	if err!= nil {
		return apiResponse(http.StatusBadRequest, ErrorBody{
			aws.String(err.Error()),
		})
	}
	return apiResponse(http.StatusOK, result)
}


//DELETE method
func DeleteUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI)(
	*events.APIGatewayProxyResponse, error,
){
	//Using delete user function from user.go 
	err := user.DeleteUser(req, tableName, dynaClient)

	if err!= nil {
		return apiResponse(http.StatusBadRequest, ErrorBody{
			aws.String(err.Error()),
		})
	}
	return apiResponse(http.StatusOK, nil)
}

//For when the request cannot be handled by the defined methods
func UnhandledMethod()(*events.APIGatewayProxyResponse, error){
	return apiResponse(http.StatusMethodNotAllowed, ErrorMethodNotAllowed)
}