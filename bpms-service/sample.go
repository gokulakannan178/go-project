package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

// Episode represents the schema for the "Episodes" collection
type Episode struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Podcast     primitive.ObjectID `bson:"podcast,omitempty"`
	Title       string             `bson:"title,omitempty"`
	Description string             `bson:"description,omitempty"`
	Duration    int32              `bson:"duration,omitempty"`
}

func main1() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27018,localhost:27019,localhost:27020/?replicaSet=rsSample"))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(context.TODO())

	database := client.Database("quickstart")
	episodesCollection := database.Collection("episodes")

	// database.RunCommand(context.TODO(), bson.D{{"create", "episodes"}})
	wc := writeconcern.New(writeconcern.WMajority())
	rc := readconcern.Snapshot()
	txnOpts := options.Transaction().SetWriteConcern(wc).SetReadConcern(rc)

	session, err := client.StartSession()
	if err != nil {
		panic(err)
	}
	defer session.EndSession(context.Background())

	err = mongo.WithSession(context.Background(), session, func(sessionContext mongo.SessionContext) error {
		if err = session.StartTransaction(txnOpts); err != nil {
			return err
		}
		result, err := episodesCollection.InsertOne(
			sessionContext,
			Episode{
				Title:    "A Transaction Episode for the Ages",
				Duration: 15,
			},
		)
		if err != nil {
			return err
		}
		fmt.Println(result.InsertedID)
		result, err = episodesCollection.InsertOne(
			sessionContext,
			Episode{
				Title:    "Transactions for All",
				Duration: 1,
			},
		)
		if err != nil {
			return err
		}
		// eps:=[]Episode{
		// 	Episode{
		// 		Title:    "one",
		// 		Duration: 15,
		// 	},
		// 	Episode{
		// 		Title:    "two",
		// 		Duration: 15,
		// 	},
		// }
		eps := []interface{}{
			// Episode{
			// 	Title:    "one",
			// 	Duration: 15,
			// },
		}
		_, err = database.Collection("two").InsertMany(
			sessionContext,
			eps,
		)
		if err != nil {
			return err
		}
		if err = session.CommitTransaction(sessionContext); err != nil {
			return err
		}
		fmt.Println(result.InsertedID)
		return nil
	})
	if err != nil {
		if abortErr := session.AbortTransaction(context.Background()); abortErr != nil {
			panic(abortErr)
		}
		panic(err)
	}
}
