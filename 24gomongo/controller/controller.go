package controller

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	repository "github.com/md-aryan-patel/gomongo/Repository"
	"github.com/md-aryan-patel/gomongo/helpers"
	"github.com/md-aryan-patel/gomongo/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionstring = "mongodb://localhost:27017/"
const dbName = "netflix"
const collName = "watchlist"

var collection *mongo.Collection

// Connect with mongodb
// Init function more line constructor
func init() {
	// client options
	clientOptions := options.Client().ApplyURI(connectionstring)

	// connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOptions)
	helpers.CheckForNilError(err, "Mongo connection successfull")
	collection = client.Database(dbName).Collection(collName)
}

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	allMovies := repository.GetAllMovies(collection)
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	repository.InsetOneMovie(movie, collection)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	params := mux.Vars(r)
	repository.UpdateOneMovie(params["id"], collection)
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	params := mux.Vars(r)
	repository.DeleteOneMovie(params["id"], collection)
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	deleteCount := repository.DeleteAllMoves(collection)
	json.NewEncoder(w).Encode(deleteCount)
}
