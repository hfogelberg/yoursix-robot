package apis

import (
	"encoding/json"
	"github.com/hfogelberg/yourSixRobot/database"
	"github.com/hfogelberg/yourSixRobot/types"
	"log"
	"math"
	"math/rand/v2"
	"net/http"
)

// RoomHandler initializes a new room with height and width.
// The new room is saved to the DB and room ID is returned
func RoomHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Printf("error in RoomHandler: called with %v, returning Bad request", r.Method)
		respondWithError(w, http.StatusMethodNotAllowed)
		return
	}

	var roomIn types.RoomSetup
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	if err := decoder.Decode(&roomIn); err != nil {
		log.Printf("error in RoomHandler body: %v", err.Error())
		respondWithError(w, http.StatusBadRequest)
		return
	}

	if err := newRoomIsValid(roomIn); err != nil {
		log.Printf("error in RoomHandler: %v", err)
		respondWithError(w, http.StatusNotAcceptable)
		return
	}

	// Generate a unique random value to identify the roomIn
	roomID := rand.IntN(math.MaxUint32)
	newRoom := types.Room{
		ID:     roomID,
		Width:  roomIn.Width,
		Height: roomIn.Height,
	}

	if err := database.SaveRoom(newRoom); err != nil {
		log.Printf("error in RoomHandler: %v", err)
		respondWithError(w, http.StatusInternalServerError)
		return
	}

	// We need to return the room ID to be used with subsequent requests
	roomOut := types.RoomOut{ID: roomID}
	respondWithJSON(w, roomOut)
}

// StartRobotHandler positions the robot in the room, facing a directions
func StartRobotHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		log.Printf("error in StartRobotHandler: called with %v, returning Bad request", r.Method)
		respondWithError(w, http.StatusMethodNotAllowed)
		return
	}

	var robotIn types.StartRobotIn
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	if err := decoder.Decode(&robotIn); err != nil {
		log.Printf("error in RoomHandler body: %v", err.Error())
		respondWithError(w, http.StatusBadRequest)
		return
	}

	if err := robotInputIsValid(robotIn); err != nil {
		log.Printf("error in StartRobotHandler: %v", err)
		respondWithError(w, http.StatusNotAcceptable)
		return
	}

	robot := types.Robot{
		RoomID:    robotIn.RoomID,
		XPosition: robotIn.XPosition,
		YPosition: robotIn.YPosition,
		Direction: robotIn.Direction,
	}

	room, err := database.GetRoom(robot)
	if err != nil {
		log.Printf("error in StartRobotHandler: %v", err)
		respondWithError(w, http.StatusInternalServerError)
		return
	}

	if ok := startPositionIsInBounds(room, robot); !ok {
		log.Printf("info: start position out of bounds for %+v", robotIn)
		respondWithError(w, http.StatusConflict)
		return
	}

	if err := database.ChangeRoom(robot); err != nil {
		log.Printf("error in StartRobotHandler: %v", err)
		respondWithError(w, http.StatusInternalServerError)
		return
	}

}
