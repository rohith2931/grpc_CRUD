package schema

import (
	"fmt"

	pb "example.com/grpc/empmgmt"
	"github.com/jinzhu/gorm"
)

func SetUpDb() {
	db, err := gorm.Open("postgres", "user=postgres password=root dbname=postgres sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	db.DropTable(&pb.Emp{})
	db.CreateTable(&pb.Emp{})
	emps := []pb.Emp{
		{Name: "rohith", Did: 3, Mid: 3},
		{Name: "Jhon", Did: 2, Mid: 1},
		{Name: "ram", Did: 1, Mid: 4},
	}
	for _, i := range emps {
		db.Create(&i)
	}
	fmt.Println("Schema is Setup")

}
