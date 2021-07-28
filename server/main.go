package main

import (
	db "../db"
	employee "../employee"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc"
	"log"
	"net"

	_ "go.mongodb.org/mongo-driver/bson"
	_ "go.mongodb.org/mongo-driver/mongo"
	_ "go.mongodb.org/mongo-driver/mongo/options"
	_ "go.mongodb.org/mongo-driver/mongo/readpref"
)
const (
	port = ":50051"
)
var m map[string]employee.EmployeeDetails
type server struct {
	employee.UnimplementedEmployeeServer
}

func (servr * server) GetEmployeeDetails(ctx context.Context, request *employee.EmployeeQuery) (*employee.EmployeeDetails, error) {
	//map part
	//rply := m[request.Id]
	//return &rply, nil

	//DB
	//connect
	log.Printf("In server get code for Id: %v", request.GetId())
	ctxdb, cancel, client, err := db.Connect("mongodb://localhost:27017")
	if err != nil {
		log.Fatal("error in connection")
		panic(err)
	}
	//Close
	defer db.Close(cancel, client, ctxdb)
	//try pinging
	db.CheckPing(client, ctxdb)

	//try getting from db
	queryResult := db.GetOne(client, ctxdb, "employeeTest", "employee", bson.M{"id":request.GetId()})
	fmt.Println("debug queryResult: ", queryResult)
	var result bson.M
	if queryResult!= nil {
		result = queryResult[0]
	}
	IdI := result["id"]  //get result
	Id := fmt.Sprintf("%v", IdI)	//convert to string

	return &employee.EmployeeDetails{Id: Id, Name: fmt.Sprintf("%v", result["name"]), Role: fmt.Sprintf("%v", result["role"])}, nil
	// queryResult
}

func (server)  AddEmployeeDetails(ctx context.Context, request *employee.EmployeeDetails) (*employee.EmployeeResponse, error){
	//map part
	//m[request.GetId()] = *request
	//return &employee.EmployeeResponse{Code: "InsertedTrue"}, nil

	//DB
	//connect
	log.Printf("In server insert code for Id: %v", request.GetId())
	ctxdb, cancel, client, err := db.Connect("mongodb://localhost:27017")
	if err != nil {
		log.Fatal("error in connection")
		panic(err)
	}
	//Close
	defer db.Close(cancel, client, ctxdb)
	//try pinging
	db.CheckPing(client, ctxdb)

	//try inserting
	document := bson.D{
		{"id", request.GetId()},
		{"name" , request.GetName()},
		{"role", request.GetRole()},
	}
	result, err := db.InsertOne(client, ctx, "employeeTest", "employee", document)
	return &employee.EmployeeResponse{Code: fmt.Sprintf("%v", result)}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	employee.RegisterEmployeeServer(s, &server{})
	//pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	//used map when we had no db
	//m = make(map[string]employee.EmployeeDetails)
	//m["101"] = employee.EmployeeDetails{Id: "101", Name: "Ankur", Role: "MTS"}

	//serve part
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}