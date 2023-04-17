package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
  
  func MongoConn() *mongo.Client {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb+srv://namaanekamal:9DC67HIMQTBaFTG5@cluster1.ddoifk0.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPI)
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create a new client and connect to the server
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
	  panic(err)
	}

	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	return client
  }

  func MongoCollection(coll string, client *mongo.Client) *mongo.Collection {
	return client.Database("db_penggajian").Collection(coll)
  }
  