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
	fmt.Println(user)

	av, err := dynamodbattribute.MarshalMap(user)

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("users-service-dev"),
	}
	_, err = U.DynamodbClient.PutItem(input)
	if err != nil {
		fmt.Println(err.Error())
		return nil, nil
	}
	log.Printf("We have inserted a new item!\n")
	return &NewUserResponse{Name: user.Name}, nil
}
