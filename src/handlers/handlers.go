package handlers

import (
	"dao"
	"encoding/json"
	"fmt"
	"net/http"
	"schema"

	"github.com/gorilla/mux"
)

//*********************Boards*********************//
var boards []schema.Board

// GetBoardEndpoint gets a Board (By its Id)
func GetBoardEndpoint(w http.ResponseWriter, r *http.Request) {
	ownerID := mux.Vars(r)["ownerid"]
	boardID := mux.Vars(r)["id"]

	payload := dao.GetBoard(ownerID, boardID)
	for _, b := range payload {
		json.NewEncoder(w).Encode(b)
		return
	}
	json.NewEncoder(w).Encode("Board not found")
}

// GetAllBoardsEndpoint gets all boards
func GetAllBoardsEndpoint(w http.ResponseWriter, r *http.Request) {
	ownerID := mux.Vars(r)["ownerid"]

	payload := dao.GetAllBoards(ownerID)
	json.NewEncoder(w).Encode(payload)
}

// CreateBoardEndpoint creates a Board
func CreateBoardEndpoint(w http.ResponseWriter, r *http.Request) {

	ownerID := mux.Vars(r)["ownerid"]

	var board schema.Board
	_ = json.NewDecoder(r.Body).Decode(&board)

	fmt.Println("board created", board)

	dao.InsertOneBoard(board, ownerID)
	json.NewEncoder(w).Encode(board)
}

// DeleteBoardEndpoint delets a Board
func DeleteBoardEndpoint(w http.ResponseWriter, r *http.Request) {

	ownerID := mux.Vars(r)["ownerid"]

	var board schema.Board
	_ = json.NewDecoder(r.Body).Decode(&board)

	fmt.Println("board deleted", board)

	dao.DeleteBoard(board, ownerID)

	json.NewEncoder(w).Encode(board)
}

// DeleteBoardByIDEndpoint delets a Board by its Id
func DeleteBoardByIDEndpoint(w http.ResponseWriter, r *http.Request) {

	boardID := mux.Vars(r)["id"]

	ownerID := mux.Vars(r)["ownerid"]

	fmt.Println("board deleted", boardID)

	dao.DeleteBoardByID(boardID, ownerID)

	json.NewEncoder(w).Encode("board deleted: " + boardID)

}

// UpdateBoardEndpoint updates a board
func UpdateBoardEndpoint(w http.ResponseWriter, r *http.Request) {

	boardID := mux.Vars(r)["id"]

	ownerID := mux.Vars(r)["ownerid"]

	var board schema.Board
	_ = json.NewDecoder(r.Body).Decode(&board)
	dao.UpdateBoard(board, boardID, ownerID)

	json.NewEncoder(w).Encode("job updated: " + boardID)

}

//*********************Job specific, shouldn't have to be used, more for reference*********************//

// CreateJobEndpoint creates a job, can be handled by UpdateBoardEndpoint
func CreateJobEndpoint(w http.ResponseWriter, r *http.Request) {
	boardID := mux.Vars(r)["id"]

	ownerID := mux.Vars(r)["ownerid"]

	categoryID := mux.Vars(r)["id2"]

	var board schema.Board
	_ = json.NewDecoder(r.Body).Decode(&board)
	dao.CreateJob(board, boardID, categoryID, ownerID)

	json.NewEncoder(w).Encode("Job created for: " + boardID)

}
