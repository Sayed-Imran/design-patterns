package db

import (
	"context"
	"fmt"

	"github.com/sayed-imran/go-design-pattern/config"
	"github.com/sayed-imran/go-design-pattern/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	database   = "metadata"
	collection = "users"
)

type MongodbRepo struct {
	MongodbClient *mongo.Client
}

func CreateNewMongoDbRepo(c *config.Config) *MongodbRepo {
	return &MongodbRepo{
		MongodbClient: c.Mongo,
	}
}

func (mr *MongodbRepo) AddUser(ctx context.Context, u models.User) error {
	usersCollection := mr.MongodbClient.Database(database).Collection(collection)
	insertResult, err := usersCollection.InsertOne(ctx, u)
	if err != nil {
		return err
	}
	fmt.Println(" --------------- Inserted a single User: ", insertResult.InsertedID)
	return nil
}

func (mr *MongodbRepo) AddMultipleUsers(ctx context.Context, usersTobeInserted ...models.User) error {
	usersCollection := mr.MongodbClient.Database(database).Collection(collection)

	users := []interface{}{}
	for _, user := range usersTobeInserted {
		users = append(users, user)
	}
	insertManyResult, err := usersCollection.InsertMany(ctx, users)
	if err != nil {
		return err
	}
	fmt.Println(" --------------- Inserted multiple documents: ", insertManyResult.InsertedIDs)
	return nil
}

func (mr *MongodbRepo) FindSingleUser(ctx context.Context, id string) (models.User, error) {
	usersCollection := mr.MongodbClient.Database(database).Collection(collection)

	filter := bson.D{{Key: "id", Value: id}}

	var result models.User

	err := usersCollection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return result, err
	}

	fmt.Printf(" --------------- Found a single document: %+v\n", result)
	return result, nil
}

func (mr *MongodbRepo) FindMultipleUsers(ctx context.Context, limit int64) ([]*models.User, error) {
	usersCollection := mr.MongodbClient.Database(database).Collection(collection)

	findOptions := options.Find()
	findOptions.SetLimit(limit)

	var results []*models.User
	cur, err := usersCollection.Find(ctx, bson.D{{}}, findOptions)
	if err != nil {
		return results, err
	}


	for cur.Next(ctx) {
		var elem models.User
		err := cur.Decode(&elem)
		if err != nil {
			return results, err
		}
		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		return results, err
	}

	cur.Close(ctx)

	fmt.Printf(" --------------- Found multiple documents (array of pointers): %+v\n", results) //array of pointer
	fmt.Println("- - -Here's the second found user (value of pointer 2) : ", *results[2])       // get the value of the pointers
	return results, nil
}

func (mr *MongodbRepo) UpdateUser(ctx context.Context, id string, u models.User) (models.User, error) {
	usersCollection := mr.MongodbClient.Database(database).Collection(collection)
	filter := bson.D{{Key: "id", Value: id}}

	var result models.User

	err := usersCollection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return result, err
	}

	usersCollection.FindOneAndReplace(ctx, filter, u)
	fmt.Printf(" --------------- Updated a single document: %+v\n", u)
	return u, nil

}


func (mr *MongodbRepo) DeleteUser(ctx context.Context, id string) error {
	usersCollection := mr.MongodbClient.Database(database).Collection(collection)

	deleteResult, err := usersCollection.DeleteMany(ctx, bson.D{{Key: "id", Value: id}})
	if err != nil {
		return err
	}
	fmt.Printf(" --------------- Deleted %v documents in the users collection\n", deleteResult.DeletedCount)
	return nil
}


func (mr *MongodbRepo) DeleteAllUsers(ctx context.Context) error {
	usersCollection := mr.MongodbClient.Database(database).Collection(collection)
	deleteResult, err := usersCollection.DeleteMany(ctx, bson.D{{}})
	if err != nil {
		return err
	}
	fmt.Printf(" --------------- Deleted %v documents in the users collection\n", deleteResult.DeletedCount)
	return nil
}

func (mr *MongodbRepo) DisconnectDB(ctx context.Context) error {
	err := mr.MongodbClient.Disconnect(ctx)
	if err != nil {
		return err
	}
	fmt.Println(" --------------- Connection to MongoDB has been closed.")
	return nil
}
