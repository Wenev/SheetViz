package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Fail to load .env")
	}

	db_str := os.Getenv("MYSQL_DB_CONN")

	db, err := gorm.Open(mysql.Open(db_str))
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	// init ORM
	db.AutoMigrate()

	//init GIN
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello World!"})
	})

	router.Run(":3001")
}
