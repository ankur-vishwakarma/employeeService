# employeeService
Basic functionality using golang, gRPC and protobuf

Run server/main.go on one terminal

`go run server/main/go` 

( GO111MODULE=off go run server/main.go )

On other terminal you can run client using which insertion and retrieval can be done from server

**GET**

`go run client/main.go get ID` 

(GO111MODULE=off go run client/main.go get 101)

**INSERT**

`go run client/main.go insert ID NAME ROLE`

(GO111MODULE=off go run client/main.go insert 101 Ankur MTS)
