package main

import (
	"fmt"
	"log"

	"github.com/LatihanFW/controllers"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/users", controllers.GetAllUsers)
	e.POST("/users", controllers.InsertUser)
	e.DELETE("/users/:user_id", controllers.DeleteUser)
	e.PUT("/users/:user_id", controllers.UpdateUser)

	fmt.Println("Connected to port 8080")
	log.Println("Connected to port 8080")
	log.Fatal(e.Start(":8080"))
}
