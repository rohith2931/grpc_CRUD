package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "example.com/grpc/empmgmt"
	"google.golang.org/grpc"
)

const (
	address = "localhost:9099"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewEmpManagementClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//Creating a new employee
	cre_res, err := c.CreateNewEmp(ctx, &pb.NewEmp{Name: "rohith", Mid: 1, Did: 2})
	if err != nil {
		log.Fatalf("Error getting employess %v", err)
	}
	fmt.Printf("\nCreated  Employee:\n %v", cre_res)
	//Get all employees
	get_res, err := c.GetEmps(ctx, &pb.GetEmpParams{})
	if err != nil {
		log.Fatalf("Error getting employess %v", err)
	}
	// fmt.Println(r)
	fmt.Println("\nAll Employess are ")
	for _, i := range get_res.Employees {
		fmt.Println(i)
	}
	// fmt.Printf("\nAll Emps:\n %v", get_res.GetEmployees())

	//Update an employee
	upd_res, err := c.UpdateEmp(ctx, &pb.Emp{Id: 3, Name: "raju"})
	if err != nil {
		log.Fatalf("\nError Updating employe %v", err)
	}
	fmt.Printf("\nUpdated Employee : %v", upd_res)

	//Delete a employee
	del_res, err := c.DeleteEmp(ctx, &pb.Emp{Id: 2})
	if err != nil {
		log.Fatalf("Error Deleting employe %v", err)
	}
	fmt.Printf("\nDeleted employee: %v", del_res)
}
