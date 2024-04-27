package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gpunit2417/MongoAPI/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"		
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://punitgoyal106:backendwithgo@cluster0.qadvr3k.mongodb.net/"
const dbName = "netflix"
const colName = "watchlist"

// most important database connection
var collection *mongo.Collection

func init(){
	//client connection
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongo db
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Mongodb connection successful")

	collection = client.Database(dbName).Collection(colName)

	fmt.Println("Client connection instance is ready")
}

//mongod helpers - file

func insertOneMovie(movie model.Netflix){
	insert, err := collection.InsertOne(context.Background(), movie)

	if err!=nil{
		log.Fatal(err)
	}

	fmt.Println("Movie successfully inserted with id ", insert.InsertedID)
}

func updateOneMovie(movieId string){
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id":id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err!=nil{
		log.Fatal(err)
	}

	fmt.Println("Modified count is: ", result.ModifiedCount)
}

func deleteOneMovie(movieId string){
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id":id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)

	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("The movie with the delete count: ", deleteCount)
}

// func deleteManyMovie(){
// 	deletedResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
// 	if err!=nil{
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Deleted count is: ", deletedResult.DeletedCount)
// }


// for return from the function
func deleteManyMovie() int64{
	deletedResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err!=nil{
		log.Fatal(err)
	}

	fmt.Println("Deleted count is: ", deletedResult.DeletedCount)
	return deletedResult.DeletedCount
}

func getAllMovies() []primitive.M{
	cur, err := collection.Find(context.Background(), bson.M{})

	if err!=nil{
		log.Fatal(err)
	}

	var movies []primitive.M

	for cur.Next(context.Background()){
		var movie bson.M
		err := cur.Decode(&movie)
		if err!=nil{
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}

	defer cur.Close(context.Background())
	return movies
}


//Actual controllers - file
func GetMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix

	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteManyMovie()
	json.NewEncoder(w).Encode(count)
}