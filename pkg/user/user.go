package user

import (
	"encoding/json"
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"main.go/pkg/validators"
)

var (
	ErrorFailedToFetchRecord ="failed to fetch record"
	ErrorFailedToUnmarshalRecord ="failed to unmarshal record"
	ErrorInvalidUserData = "failed to validate user data"
	ErrorInvalidEmail = "failed to validate email"
	ErrorFailedToMarshalRecord = "failed to marshal record"
	ErrorCouldNotDeleteItem = "could not delete item"
	ErrorCouldNotDynamoPutItem = "could not dynamo put item"
	ErrorUserAlreadyExists = "user.User already exists"
	ErrorUserDoesNotExist = "user.User does not exist"
)

type User struct {
	Email 		string `json:"email"`
	FirstName 	string `json:"firstName"`
	LastName 	string `json:"lastName"`
}

func FetchUser(email, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*User, error) {

	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"email":{
				s: aws.String(email)
			}
		},
		TableName: aws.String(tableName),
	}

	result, arr := dynaClient.GetItem(input)
	if err != nil {
		return nil, errors.New(ErrorFailedToFetchRecord)
	}

	item := new(User)
	err = dynamodbattribute.UnmarshalMap(result.Item, item)
	if err != nil {
		return nil, errors.New(ErrorFailedToUnmarshalRecord)
	}

	return item, nil
}

//Extra function
func FetchUsers(tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*[]User, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String(tableName)
	}
	result, err := dynaClient.Scan(input)
	if err != nil {
		return nil, errors.New(ErrorFailedToFetchRecord)
	}
	item := new([]User)
	err = dynamodbattribute.UnmarshalMap(result.Items, item)
	if err != nil {
		return nil, errors.New(ErrorFailedToUnmarshalRecord)
	}

	return item, nil
}

func CreateUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*User, error) {
	var u User

	if err := json.UnmarshalMap([]byte(req.body), &u); err != nil {
		return nil, errors.New(ErrorInvalidUserData)
	}

	if !validators.IsEmailValid(u.Email) {
		return nil, errors.New(ErrorInvalidEmail)
	}

	currentUser, _ := FetchUser(u.Email, tableName, dynaClient)
	if currentUser != nil && len(currentUser.Email) != 0 {
		return nil, errors.New(ErrorUserAlreadyExists)
	}

	av, err := dynamodbattribute.MarshalMap(u)
	if err != nil {
		return nil, errors.New(ErrorFailedToMarshalRecord)
	}

	input := &dynamodb.PutItemInput{
		Item: av,
		tableName: aws.String(tableName)
	} 

	_, err := dynaClient.PutItem(input)
	if err != nil {
		return nil, errors.New(ErrorCouldNotDynamoPutItem)
	}
}

func UpdateUser() () {

}

func DeleteUser() error {

}
