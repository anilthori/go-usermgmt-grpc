package usermgmt

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	database "github.com/anilthori/go-usermgmt-grpc/redis"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	// "github.com/go-redis/redis"
)

type UserServer struct {
	DynamodbClient *dynamodb.DynamoDB
	RedisClient    database.Database
	TableName      string
}

func (U *UserServer) CreateNewUser(ctx context.Context, user *NewUserRequest) (*NewUserResponse, error) {
	// storing in DynamoDB
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
	suser, _ := json.Marshal(user)
	U.RedisClient.Set(user.UserId, string(suser))

	log.Printf("We have inserted a new item!\n")
	return &NewUserResponse{UserId: user.UserId}, nil
}

func (U *UserServer) GetUser(ctx context.Context, message *UserDetailsRequest) (*UserDetailsResponse, error) {
	val, _ := U.RedisClient.Get(message.UserId)

	if val != "" {
		return &UserDetailsResponse{
			UserId: val,
		}, nil
	}

	return &UserDetailsResponse{
		UserId: message.UserId,
	}, nil

}
