package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	model "github.com/shahanahmed86/mongo-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString string = "mongodb://admin:lmelg8@localhost:27017/netflix"
const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

// connect with mongodb

func init() {
	// client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mongo DB connection success")

	collection = client.Database(dbName).Collection(colName)

	// collection instance
	fmt.Println("Collection instance is ready")
}

// actual controller - file
func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	allMovies := getMovies()

	json.NewEncoder(w).Encode(allMovies)
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")

	params := mux.Vars(r)

	movie := getMovie(params["id"])

	json.NewEncoder(w).Encode(movie)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		log.Fatal(err)
	}

	createMovie(movie)

	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)

	updateMovie(params["id"])

	json.NewEncoder(w).Encode("The movie marked as watched")
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)

	deleteMovie(params["id"])

	json.NewEncoder(w).Encode("The movie has been deleted")
}

func DeleteMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	deleteMovies()

	json.NewEncoder(w).Encode("All movies has been deleted")
}

// mongodb helpers - file

func createMovie(movie model.Netflix) {
	result, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted one movie in DB with id", result.InsertedID)
}

func updateMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, updateErr := collection.UpdateOne(context.Background(), filter, update)
	if updateErr != nil {
		log.Fatal(updateErr)
	}

	fmt.Println("modified count", result.ModifiedCount)
}

func deleteMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}

	result, deleteErr := collection.DeleteOne(context.Background(), filter)
	if deleteErr != nil {
		log.Fatal(deleteErr)
	}

	fmt.Println("deleted count", result.DeletedCount)
}

func deleteMovies() {
	result, err := collection.DeleteMany(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted count", result.DeletedCount)
}

func getMovies() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M

	for cursor.Next(context.Background()) {
		var movie bson.M
		curErr := cursor.Decode(&movie)
		if curErr != nil {
			log.Fatal(curErr)
		}
		movies = append(movies, movie)
	}

	defer cursor.Close(context.Background())
	return movies
}

func getMovie(movieId string) model.Netflix {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}

	result := collection.FindOne(context.Background(), filter)

	var movie model.Netflix
	findErr := result.Decode(&movie)
	if findErr != nil {
		log.Fatal(findErr)
	}
	return movie
}
