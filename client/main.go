package main

import (
	"../employee"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"os"
	"time"
)

const(
	address     = "localhost:50051"
)

func main(){
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	clnt := employee.NewEmployeeClient(conn)

	//try getting the args here
	id := "101"  //added by server just for example
	var name, role, operation string
	if len(os.Args) > 1 {
		operation = os.Args[1]
		id = os.Args[2]
		if operation == "insert"{
			name = os.Args[3]
			role = os.Args[4]
		}
	}
	fmt.Println("DEBUG: " + operation + " " + id + " " + name + " " + role)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if operation == "insert" {
		resp, err := clnt.AddEmployeeDetails(ctx, &employee.EmployeeDetails{Id: id, Name: name, Role: role})
		if err != nil {
			log.Fatalf("could not get reponse: %v", err)
		}
		log.Printf("Response is: %s", resp)
	} else if operation == "get" {
		resp, err := clnt.GetEmployeeDetails(ctx, &employee.EmployeeQuery{Id: id})
		if nil != err {
			log.Fatalf("could not get reponse: %v", err)
		}
		log.Printf("\nResponse is: %s", resp)
	} else {
		resp, err := clnt.GetEmployeeDetails(ctx, &employee.EmployeeQuery{Id: id}) //was added by server
		if err != nil {
			log.Fatalf("could not get reponse: %v", err)
		}
		log.Printf("\nResponse is: %s", resp)
	}


}