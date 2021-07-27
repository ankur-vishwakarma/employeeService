package main

import (
	employee "../employee"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)
const (
	port = ":50051"
)
var m map[string]employee.EmployeeDetails
type server struct {
	employee.UnimplementedEmployeeServer
}

func (servr * server) GetEmployeeDetails(ctx context.Context, request *employee.EmployeeQuery) (*employee.EmployeeDetails, error) {
	rply := m[request.Id]
	return &rply, nil
}

func (server)  AddEmployeeDetails(ctx context.Context, request *employee.EmployeeDetails) (*employee.EmployeeResponse, error){
	m[request.GetId()] = *request
	return &employee.EmployeeResponse{Code: "InsertedTrue"}, nil
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
	m = make(map[string]employee.EmployeeDetails)
	m["101"] = employee.EmployeeDetails{Id: "101", Name: "Ankur", Role: "MTS"}
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}