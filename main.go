package main

import (
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", GetUser)
	e.POST("/", CreateUserPost)
	e.PUT("/", UpdateUser)
	e.DELETE("/", DeleteUser)

	// Start server
	e.Logger.Fatal(e.Start(":80"))

	// if err := e.StartTLS(":8443", "server.crt", "server.key"); err != http.ErrServerClosed {
	// 	log.Fatal(err)
	// }
}

type User struct {
	ID			primitive.ObjectID	`bson:"_id,omitempty" json:"_id,omitempty"`
	Name 		string				`bson:"name" json:"name"`
	Age 		string				`bson:"age" json:"age"`
	CreatedAt 	time.Time			`bson:"created_at" json:"created_at"`
}

