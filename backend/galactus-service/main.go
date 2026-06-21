// package main

// import "fmt"

// // = examples (package-level)
// var x string = "Hello World!"
// var y int = 10
// var z bool = true

// const pi = 3.14159

// func main() {
// 	a := 10
// 	b := 30
// 	fmt.Printf("BABA\nBABA\n%f", pi)
// 	fmt.Println(x, y, z)
// 	fmt.Println(a, b)
// }

package main

import (
	"context"
	"log"
	"net"
	"os"

	pb "github.com/Wenev/SheetViz/backend/contracts/galactus-service/gen"
	"github.com/Wenev/SheetViz/backend/galactus-service/controller"
	"github.com/Wenev/SheetViz/backend/galactus-service/model"
	"github.com/Wenev/SheetViz/backend/galactus-service/repository"
	"github.com/Wenev/SheetViz/backend/galactus-service/service"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type GalactusServer struct {
	pb.UnimplementedGalactusServiceServer
	c *controller.UserController
}

func NewGalactusServer(c *controller.UserController) *GalactusServer {
	return &GalactusServer{c: c}
}

func (s *GalactusServer) GetUserByID(context.Context, *pb.GetUserByIDRequest) (*pb.UserResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method GetUserByID not implemented")

}
func (s *GalactusServer) GetUsers(context.Context, *emptypb.Empty) (*pb.GetUsersResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method GetUsers not implemented")
}
func (s *GalactusServer) UpdateUser(context.Context, *pb.UpdateUserRequest) (*pb.UserResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method UpdateUser not implemented")
}
func (s *GalactusServer) DeleteUser(context.Context, *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method DeleteUser not implemented")
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	db_string := os.Getenv("MYSQL_CONN")

	db, err := gorm.Open(mysql.Open(db_string))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.User{})

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatalf("failed to listen")
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGalactusServiceServer(grpcServer, NewGalactusServer(userController))

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal("failed to serve grPc")
	}

	// r := gin.Default()

	// r.GET("/hello", func(c *gin.Context) {
	// 	c.String(http.StatusOK, "Hello, World!")
	// })

	// v1 := r.Group("/api/v1")
	// {
	// 	userController.RegisterRoutes(v1)
	// }

	// r.Run(":3001")
}
