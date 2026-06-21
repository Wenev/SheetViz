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
	"net/http"
	"os"

	"github.com/Wenev/SheetViz/backend/galactus-service/controller"
	"github.com/Wenev/SheetViz/backend/galactus-service/model"
	"github.com/Wenev/SheetViz/backend/galactus-service/repository"
	"github.com/Wenev/SheetViz/backend/galactus-service/service"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

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

	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	v1 := r.Group("/api/v1")
	{
		userController.RegisterRoutes(v1)
	}

	r.Run(":3001")
}	
