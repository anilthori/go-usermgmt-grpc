package main

import (
	"fmt"
	"log"
	"net"

	// "encoding/json"
	database "github.com/anilthori/go-usermgmt-grpc/redis"
	"github.com/anilthori/go-usermgmt-grpc/usermgmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"google.golang.org/grpc"
)

const (
	port = ":9000"
)

func ConnectDynamoDB() *dynamodb.DynamoDB {
	cfg := &aws.Config{
		Region:   aws.String("ap-southeast-1"),
		Endpoint: aws.String("http://localhost:8000"),
	}

	sess, err := session.NewSession(cfg)
	if err != nil {
		fmt.Println("Unable to connect to DynamoDB!")
	}

	return dynamodb.New(sess)
}

var DB *dynamodb.DynamoDB

var redisDB database.Database

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	DB := ConnectDynamoDB()
	redisDB, rerr := database.CreateRedisDatabase()

	if rerr != nil {
		fmt.Println("Unable to connect to Redis")
	}

	umgmt := usermgmt.UserServer{DB, redisDB, "users"}

	s := grpc.NewServer()

	usermgmt.RegisterUserManagementServer(s, &umgmt)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
