package main

import (
	"context"
	"log"
	"net"

	pb "example.com/grpc/empmgmt"
	"example.com/grpc/schema"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

const (
	port = ":9099"
)

type EmpManagementServer struct {
	pb.UnimplementedEmpManagementServer
	Db *gorm.DB
}

func NewEmpManagementServer() *EmpManagementServer {
	return &EmpManagementServer{}
}
func (server *EmpManagementServer) Run() error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterEmpManagementServer(s, server)
	log.Printf("Server listening at %v", lis.Addr())
	return s.Serve(lis)
}
func (s *EmpManagementServer) CreateNewEmp(ctx context.Context, in *pb.NewEmp) (*pb.Emp, error) {
	var i pb.Emp = pb.Emp{Name: in.Name, Did: in.Did, Mid: in.Mid}

	s.Db.Create(&i)
	// s.Db.Model(&pb.Emp{}).First(&i)
	// log.Printf("%+v", i)
	return &i, nil
}
func (s *EmpManagementServer) GetEmps(ctx context.Context, in *pb.GetEmpParams) (*pb.EmpList, error) {
	var res pb.EmpList
	var temp []pb.Emp
	row, _ := s.Db.Model(&pb.Emp{}).Find(&temp).Rows()
	defer row.Close()
	// fmt.Printf("%+v", temp)
	for row.Next() {
		emp := pb.Emp{}
		err := row.Scan(&emp.Id, &emp.Name, &emp.Mid, &emp.Did)
		if err != nil {
			log.Fatalf("Error in scanning %v", err)
		}
		res.Employees = append(res.Employees, &emp)
	}
	// for _,i := range temp{
	// 	res=append(res,i)
	// }
	return &res, nil
}

func (s *EmpManagementServer) UpdateEmp(ctx context.Context, in *pb.Emp) (*pb.Emp, error) {
	var res pb.Emp
	// log.Fatalf("\n\n %v", res)
	// if in.Name != "" {
	// 	res.Name = in.Name
	// }
	// if in.Did != 0 {
	// 	res.Did = in.Did
	// }
	// if in.Mid != 0 {
	// 	res.Mid = in.Mid
	// }
	// s.Db.Save(&res)
	s.Db.Model(&pb.Emp{}).Where("id=?", in.Id).Updates(in)
	s.Db.Find(&res, in.Id)

	return &res, nil
}
func (s *EmpManagementServer) DeleteEmp(ctx context.Context, in *pb.Emp) (*pb.Emp, error) {
	var res pb.Emp
	s.Db.Find(&res, in.Id)
	s.Db.Model(&pb.Emp{}).Delete(&res, in.Id)

	return &res, nil
}

func main() {
	db, err := gorm.Open("postgres", "user=postgres password=root dbname=postgres sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	schema.SetUpDb()
	var emp_mgmt_server *EmpManagementServer = NewEmpManagementServer()
	emp_mgmt_server.Db = db
	if err := emp_mgmt_server.Run(); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
