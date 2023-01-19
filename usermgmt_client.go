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

	log.Println("Enter candid Id: ")
	var user_id string
	fmt.Scan(&user_id)
	log.Println(("Enter your name:"))
	var name string
	fmt.Scan(&name)
	log.Printf("Enter your age")
	var age int32
	fmt.Scan(&age)

	message := usermgmt.NewUserRequest{
		UserId: user_id,
		Name:   name,

		Age: age,
	}
	// fmt.Println(message.Name)
	response, err := c.CreateNewUser(context.TODO(), &message)
	if err != nil {
		log.Fatalf("Error when calling PutUser: %s", err)
	}

	log.Printf("Response from Server: %s", response)

	log.Printf("Enter candid id to search:")
	var user_id1 string
	fmt.Scan(&user_id1)

	message1 := usermgmt.UserDetailsRequest{
		UserId: user_id1,
	}
	response1, err1 := c.GetUser(context.TODO(), &message1)
	if err1 != nil {
		log.Fatalf("Error when calling PutUser: %s", err1)
	}

	log.Printf("Response from Server: %s", response1)

}
