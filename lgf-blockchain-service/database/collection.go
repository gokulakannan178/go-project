package database

import (
	"context"
	"fmt"
	"os"
	"reflect"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB(name string) *mongo.Collection {
	cilentOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	fmt.Println("Clientoptions TYPE:", reflect.TypeOf(cilentOptions), "\n")

	client, err := mongo.Connect(context.TODO(), cilentOptions)
	if err != nil {
		fmt.Println("Mongo.Connect() ERROR:", err)
		os.Exit(1)
	}
	// ctx, _ := context.WithTimeout(context.Background(), 25*time.Second)

	fmt.Println("Connected to MongoDB!")

	collection := client.Database("gokul").Collection(name)

	fmt.Println("Collection type:", reflect.TypeOf(collection), "\n")

	return collection
}
func ConnectMarksheet() *mongo.Collection {
	cilentOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	fmt.Println("Clientoptions TYPE:", reflect.TypeOf(cilentOptions), "\n")

	client, err := mongo.Connect(context.TODO(), cilentOptions)
	if err != nil {
		fmt.Println("Mongo.Connect() ERROR:", err)
		os.Exit(1)
	}
	// ctx, _ := context.WithTimeout(context.Background(), 25*time.Second)

	fmt.Println("Connected to MongoDB!")

	collection := client.Database("gokul").Collection("marksheet")

	fmt.Println("Collection type:", reflect.TypeOf(collection), "\n")

	return collection
}
