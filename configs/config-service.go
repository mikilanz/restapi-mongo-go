package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func EnvMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("MONGOURI")
}

// use godot package to load/read the .env file and
// return the value of the key
func GetEnvKey(key string) string {
	_, err := os.Open(".env")
	if err == nil {
		// load .env file
		err := godotenv.Load(".env")

		if err != nil {
			fmt.Printf("Error loading .env file")
		}
	}

	return os.Getenv(key)
}

// use godot package to load/read the .env file and
// return the value of the key
func SetEnvKey(key string, value string) error {
	_, err := os.Open(".env")
	if err == nil {
		// load .env file
		err := godotenv.Load(".env")

		if err != nil {
			fmt.Printf("Error loading .env file")
		}
	}

	return os.Setenv(key, value)
}

func ConnectDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer cancel()
	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	return client
}

//Client instance
var DB *mongo.Client = ConnectDB()

//getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("golangAPI").Collection(collectionName)
	return collection
}
