package handlers

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

//Function for converting our response into json and sending to API gateway in AWS
func apiResponse(status int, body interface{}) (*events.APIGatewayProxyResponse, error) {
	resp := events.APIGatewayProxyResponse{Headers: map[string]string{"Content-Type":"application/json"}}
	resp.StatusCode = status

	//Turning to json
	stringBody, _ := json.Marshal(body)
	resp.Body = string(stringBody)
	return &resp, nil
}