package main

import (
	"ginsample/auth"
	"ginsample/handler"
	"ginsample/user"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	//dsn := "root:@tcp(127.0.0.1:3307)/db_sosmedes?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := goDotEnvVariable("DB_HOST")

	log.Println("dsn", dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService()

	token, err := authService.ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo0fQ.IkHa4TmZJDS-ZSm2Ow6oB1lYJE19NAcwfWr14N15NVk")

	if err != nil {
		log.Println("error : ", err)
	} else {
		log.Println("sukses : ", token)

	}

	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()
	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailibility)
	api.POST("/upload_avatar", userHandler.UploadAvatar)

	router.Run(":3030")

}
