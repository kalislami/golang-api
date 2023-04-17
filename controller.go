package main

import (
	"context"
	"fmt"
	"golang-api/db"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUser(c echo.Context) error {
	age := c.QueryParam("age")

	client :=db.MongoConn()
	defer client.Disconnect(context.TODO())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var users = make([]User, 0)

	col := db.MongoCollection("user", client)

	filter := bson.D{}

	println("age: ", age)

	if age != "" {
		filter = bson.D{{"age", age}}	
	}

	result, err := col.Find(ctx, filter)

	if err != nil {
		return err
	}

	defer result.Close(ctx)

	fmt.Println(users)

	for result.Next(ctx) {
		var user User
		err = result.Decode(&user)
		if err != nil {
			return err
		}
		users = append(users, user)
	}

	return c.JSON(http.StatusOK, users)
}

func CreateUserPost(c echo.Context) error {
	
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}

	client :=db.MongoConn()
	defer client.Disconnect(context.TODO())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var table = db.MongoCollection("user", client)

	_, err := table.InsertOne(ctx, User{
		Name: u.Name,
		Age: u.Age,
		CreatedAt: time.Now(),
	})

	if err != nil {
		fmt.Println("error insert user: ")
		fmt.Println(err)
	}

	return c.JSON(http.StatusCreated, u)
}

func UpdateUser(c echo.Context) error {
	
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}

	client :=db.MongoConn()
	defer client.Disconnect(context.TODO())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var table = db.MongoCollection("user", client)

	filter := bson.D{{"_id", u.ID}}
	update := bson.D{{"$set", bson.D{
		{"name", u.Name},
		{"age", u.Age},
	}}}

	_, err := table.UpdateOne(ctx, filter, update)

	if err != nil {
		fmt.Println("error update user: ")
		fmt.Println(err)
	}

	return c.JSON(http.StatusOK, "update sukses")
}

func DeleteUser(c echo.Context) error {
	
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}

	client :=db.MongoConn()
	defer client.Disconnect(context.TODO())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var table = db.MongoCollection("user", client)

	filter := bson.D{{"_id", u.ID}}

	_, err := table.DeleteOne(ctx, filter)

	if err != nil {
		fmt.Println("error update user: ")
		fmt.Println(err)
	}

	return c.JSON(http.StatusOK, "update sukses")
}