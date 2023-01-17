package main

import (
	// "context"
	"fmt"
	"log"

	// "github.com/anilthori/go-usermgmt-grpc"
	"github.com/anilthori/go-usermgmt-grpc/usermgmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:9000"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	c := usermgmt.NewUserManagementClient(conn)

	log.Println(("Enter your name:"))
	var name string
	fmt.Scan(&name)
	log.Printf("Enter your age")
	var age int64
	fmt.Scan(&age)

	message := c.NewUserRequest{
		Name: name,
		Age:  age,
	}

	response, err := c.CreateNewUser(context.TODO(), &message)
	if err != nil {
		log.Fatalf("Error when calling PutUser: %s", err)
	}

	log.Printf("Response from Server: %s", response)
}

//
