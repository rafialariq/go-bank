package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rafialariq/go-bank/controller"
	"github.com/rafialariq/go-bank/middleware"
	"github.com/rafialariq/go-bank/model"
	"github.com/rafialariq/go-bank/repository"
	"github.com/rafialariq/go-bank/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// connect database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", "localhost", "postgres", "postgres", "digitalbank", 5432, "disable")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		// logging here
		log.Fatal(err)
	}

	err = db.AutoMigrate(&model.User{}, &model.Merchant{})
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	r.Use(middleware.LogMiddleware())
	r.Use(gin.Recovery())

	loginRepo := repository.NewLoginRepository(db)
	loginService := service.NewLoginService(loginRepo)
	loginController := controller.NewLoginController(loginService)

	registerRepo := repository.NewRegisterRepository(db)
	registerService := service.NewRegisterService(registerRepo)
	registerController := controller.NewRegisterController(registerService)

	r.GET("/login", loginController.LoginHandler)
	r.POST("/signup", registerController.RegisterHandler)

	r.Run(":8080")

}
