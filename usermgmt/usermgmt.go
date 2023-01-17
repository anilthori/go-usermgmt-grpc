package usermgmt

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type UserServer struct {
	DynamodbClient *dynamodb.DynamoDB
	TableName      string
}

func (U *UserServer) CreateNewUser(ctx context.Context, user *NewUserRequest) (*NewUserResponse, error) {
	fmt.Println("hello")
	// fmt.Println(user.Name)

	av, err := dynamodbattribute.MarshalMap(user)

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("users"),
	}
	_, err = U.DynamodbClient.PutItem(input)
	if err != nil {
		fmt.Println(err.Error())
		return nil, nil
	}
	log.Printf("We have inserted a new item!\n")
	return &NewUserResponse{UserId: user.UserId}, nil
}

func (U *UserServer) GetUser(ctx context.Context, message *UserDetailsRequest) (*UserDetailsResponse, error) {
	// _, err := U.DynamodbClient.GetItem(message.)
	return &UserDetailsResponse{
		UserId: message.UserId,
		Name:   message.Name,
		Age:    message.Age,
	}, nil

}
