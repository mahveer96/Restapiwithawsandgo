package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func login(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"message": "Login successful"}
	json.NewEncoder(w).Encode(response)
}

func main() {
	//////////////////////////////////////////////////////////
	uri := "mongodb://mahesh:MyPass123@13.60.42.229:27017"

	// Set client options
	clientOpts := options.Client().ApplyURI(uri)

	// Connect with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatalf("❌ Connection error: %v", err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatalf("❌ Disconnect error: %v", err)
		}
	}()

	// Ping MongoDB
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("❌ Ping failed: %v", err)
	}
	fmt.Println("✅ Connected to MongoDB on AWS EC2!")

	// Example: Use a collection
	collection := client.Database("testdb").Collection("users")

	// Insert a test document
	doc := map[string]string{"name": "Mahesh", "role": "Developer"}
	_, err = collection.InsertOne(ctx, doc)
	if err != nil {
		log.Fatalf("❌ Insert failed: %v", err)
	}

	fmt.Println("✅ Inserted a document successfully!")

	http.HandleFunc("/login", login)
	err = http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
