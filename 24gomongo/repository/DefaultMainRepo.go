package repository

import (
	"context"
	"fmt"

	"github.com/md-aryan-patel/gomongo/helpers"
	"github.com/md-aryan-patel/gomongo/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsetOneMovie(movie model.Netflix, c *mongo.Collection) {
	inserted, err := c.InsertOne(context.Background(), movie, nil)
	helpers.CheckForNilError(err, "")
	fmt.Println("Inserted one movie with insert ID: ", inserted.InsertedID)
}

func UpdateOneMovie(movieId string, c *mongo.Collection) {
	id, err := primitive.ObjectIDFromHex(movieId)
	helpers.CheckForNilError(err, "")
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"Watched": true}}
	updated, uerr := c.UpdateOne(context.Background(), filter, update)
	helpers.CheckForNilError(uerr, "")
	fmt.Println("Modified count: ", updated.ModifiedCount)
}

func DeleteOneMovie(movieId string, c *mongo.Collection) {
	id, err := primitive.ObjectIDFromHex(movieId)
	helpers.CheckForNilError(err, "")
	filter := bson.M{"_id": id}
	deleteCount, dErr := c.DeleteOne(context.Background(), filter)
	helpers.CheckForNilError(dErr, "")
	fmt.Println("Delete count: ", deleteCount)
}

func DeleteAllMoves(c *mongo.Collection) int64 {
	deleted, err := c.DeleteMany(context.Background(), bson.D{{}}, nil)
	helpers.CheckForNilError(err, "")
	fmt.Println("Delete count: ", deleted.DeletedCount)
	return deleted.DeletedCount
}

func GetAllMovies(c *mongo.Collection) []primitive.M {
	curser, err := c.Find(context.Background(), bson.D{{}})
	helpers.CheckForNilError(err, "")
	defer curser.Close(context.Background())
	var movies []primitive.M

	for curser.Next(context.TODO()) {
		var movie bson.M
		err := curser.Decode(&movie)
		helpers.CheckForNilError(err, "")
		movies = append(movies, movie)
	}
	return movies
}
