package main

import (
	"common"
	"log"
	"net/http"
	"routers"

	"github.com/gorilla/handlers"
	// below used for init func to populate db
	// "context"
	// "encoding/json"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
	// "io/ioutil"
	// "schema"
	// "time"
)

func init() {

	// Below: Populates database with dummy data

	// // CONNECTIONSTRING DB connection string
	// const CONNECTIONSTRING = "mongodb+srv://conan:D7zlVg4sPf7v8hBw@cluster0-ctwix.mongodb.net/test?retryWrites=true"

	// // DBNAME Database name
	// const DBNAME = "test"

	// // COLLNAME Collection name
	// const COLLNAME = "boards"

	// var boards []schema.Board
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// client, err := mongo.Connect(ctx, options.Client().ApplyURI(CONNECTIONSTRING))
	// if err != nil {
	// 	log.Fatal("[Init 1]: %s\n", err)
	// }
	// db := client.Database(DBNAME)

	// // Load values from JSON file to model
	// byteValues, err := ioutil.ReadFile("data/board_data.json")
	// if err != nil {
	// 	log.Fatal("[Init 2]: %s\n", err)
	// }
	// json.Unmarshal(byteValues, &boards)

	// // Insert people into DB
	// var bds []interface{}
	// for _, b := range boards {
	// 	bds = append(bds, b)
	// }
	// _, err = db.Collection(COLLNAME).InsertMany(context.Background(), bds)
	// if err != nil {
	// 	log.Fatal("[Init 3]: %s\n", err)

	// }
}

// Entry point for the program
func main() {

	// Calls startup logic
	common.StartUp()
	// Get the mux router object
	router := routers.InitRoutes()

	// allows frontend to call api
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	log.Println("Listening...")
	http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(router))
}
