package dao

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"schema"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CONNECTIONSTRING DB connection string
const CONNECTIONSTRING = "mongodb+srv://user:8b9ax1puyPu8RUzQ@cluster0-ctwix.mongodb.net/test?retryWrites=true"

// DBNAME Database name
const DBNAME = "test"

// COLLNAME Collection name
const COLLNAME = "boards"

var db *mongo.Database

// Connect establish a connection to database
func init() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(CONNECTIONSTRING))

	if err != nil {
		log.Fatal("[init]: %s\n", err)
	}
	// Collection types can be used to access the database
	db = client.Database(DBNAME)

}

// InsertOneBoard inserts one item from board model
func InsertOneBoard(board schema.Board, ownerID string) {
	fmt.Println(board)
	board.ID = primitive.NewObjectID()
	board.OwnerID = ownerID

	_, err := db.Collection(COLLNAME).InsertOne(context.Background(), board)
	if err != nil {
		log.Fatal(err)
	}
}

// GetAllBoards returns all boards (or one board if specified) from DB
func GetAllBoards(ownerID string) []schema.Board {
	cur, err := db.Collection(COLLNAME).Find(context.Background(), bson.M{"owner_id": bson.M{"$eq": ownerID}})
	if err != nil {
		log.Fatal("[GetAllBoards 1]: %s\n", err)
	}
	var elements []schema.Board
	var elem schema.Board
	// Get the next result from the cursor
	for cur.Next(context.Background()) {
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal("[GetAllBoards 2]: %s\n", err)
		}
		elements = append(elements, elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatal("[GetAllBoards 3]: %s\n", err)
	}
	cur.Close(context.Background())
	fmt.Println("elems: ", elements)
	return elements
}

// GetBoard returns a board by owner and board id from DB
func GetBoard(ownerID string, boardID string) []schema.Board {
	idPrimitive, err := primitive.ObjectIDFromHex(boardID)
	cur, err := db.Collection(COLLNAME).Find(context.Background(), bson.M{"_id": idPrimitive, "owner_id": bson.M{"$eq": ownerID}})
	if err != nil {
		log.Fatal("[GetBoard 1]: %s\n", err)
	}
	var elements []schema.Board
	var elem schema.Board
	// Get the next result from the cursor
	for cur.Next(context.Background()) {
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal("[GetBoard 2]: %s\n", err)
		}
		elements = append(elements, elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatal("[GetBoard 3]: %s\n", err)
	}
	cur.Close(context.Background())
	fmt.Println("elems: ", elements)
	return elements
}

// DeleteBoard deletes an existing board
func DeleteBoard(board schema.Board, ownerID string) {
	board.OwnerID = ownerID
	result, err := db.Collection(COLLNAME).DeleteOne(context.Background(), board, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("delete func", result)
}

// DeleteBoardByID deletes an existing board by id
func DeleteBoardByID(id string, ownerID string) {

	// // Declare a primitive ObjectID from a hexadecimal string
	idPrimitive, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal("primitive.ObjectIDFromHex ERROR:", err)
	} else {
		// Call the DeleteOne() method by passing BSON
		res, err := db.Collection(COLLNAME).DeleteOne(context.Background(), bson.M{"_id": idPrimitive, "owner_id": bson.M{"$eq": ownerID}})
		fmt.Println("DeleteOne Result TYPE:", reflect.TypeOf(res))

		if err != nil {
			log.Fatal("DeleteOne() ERROR:", err)
		}
	}
}

// UpdateBoard updates an existing Board
func UpdateBoard(board schema.Board, boardID string, ownerID string) {

	objID, err := primitive.ObjectIDFromHex(boardID)
	if err != nil {
		log.Fatal("UpdateBoard() ERROR:", err)
	}

	filter := bson.M{"_id": bson.M{"$eq": objID}, "owner_id": bson.M{"$eq": ownerID}}

	if len(board.Categories) > 0 {
		//there is some category data to add
		UpdateCategories(filter, board)
	} else {
		fmt.Println("Only edited Title")
		UpdateTitle(filter, board)
	}

}

func UpdateCategories(filter bson.M, board schema.Board) {

	update := bson.M{"$set": bson.M{"categories": board.Categories}}

	updateResult, err := db.Collection(COLLNAME).UpdateOne(
		context.Background(),
		filter,
		update,
	)
	if err != nil {
		log.Fatal("UpdateBoard() ERROR:", err)
	}

	fmt.Println("edit job func input: ", board.Categories)

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	fmt.Println("Full Result: ", updateResult)

}
func UpdateTitle(filter bson.M, board schema.Board) {

	update := bson.M{"$set": bson.M{"title": board.Title}}

	updateResult, err := db.Collection(COLLNAME).UpdateOne(
		context.Background(),
		filter,
		update,
	)
	if err != nil {
		log.Fatal("UpdateBoard() ERROR:", err)
	}

	fmt.Println("edit job func input: ", board.Categories)

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	fmt.Println("Full Result: ", updateResult)

}

// CreateJob updates an existing Board, can be handled by UpdateBoard
func CreateJob(board schema.Board, boardID string, categoryID string, ownerID string) {

	objID, err := primitive.ObjectIDFromHex(boardID)
	categoryID_int, err := strconv.Atoi(categoryID)
	if err != nil {
		// handle error
		log.Fatal("CreateJob() str conv ERROR:", err)
	}

	//get job details from board document inputted through put request
	jd := schema.JobDetails{
		Company:     board.Categories[0].Jobs[0].JobDetails.Company,
		Title:       board.Categories[0].Jobs[0].JobDetails.Title,
		Location:    board.Categories[0].Jobs[0].JobDetails.Location,
		Category:    board.Categories[0].Jobs[0].JobDetails.Category,
		PostDate:    board.Categories[0].Jobs[0].JobDetails.PostDate,
		Description: board.Categories[0].Jobs[0].JobDetails.Description,
		Experience:  board.Categories[0].Jobs[0].JobDetails.Experience,
		URL:         board.Categories[0].Jobs[0].JobDetails.URL,
		DateAdded:   board.Categories[0].Jobs[0].JobDetails.DateAdded,
		Salary:      board.Categories[0].Jobs[0].JobDetails.Salary,
		Tasks:       board.Categories[0].Jobs[0].JobDetails.Tasks}

	//stick job details into job var
	j := schema.Job{JobDetails: jd, ID: board.Categories[0].Jobs[0].ID}

	filter := bson.M{"_id": bson.M{"$eq": objID}, "categories.id": bson.M{"$eq": categoryID_int}}
	update := bson.M{"$set": bson.M{"categories.$.jobs": j}}

	updateResult, err := db.Collection(COLLNAME).UpdateOne(
		context.Background(),
		filter,
		update,
	)
	if err != nil {
		log.Fatal("CreateJob() ERROR:", err)
	}

	fmt.Println("create job func input: ", j)

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	fmt.Println("Full Result: ", updateResult)

}

// InsertManyValues inserts many items from byte slice
// func InsertManyValues(people []schema.Board) {
// 	var ppl []interface{}
// 	for _, p := range people {
// 		ppl = append(ppl, p)
// 	}
// 	_, err := db.Collection(COLLNAME).InsertMany(context.Background(), ppl)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
