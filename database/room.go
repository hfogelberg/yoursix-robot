package database

import (
	"encoding/json"
	"fmt"
	"github.com/hfogelberg/yourSixRobot/types"
	"os"
)

const dbFiles = "./dbfiles/"

// Note 1: In a real world application I would have used an RDS here,
// but cannot since the driver is not in the standard package
// Note 2: Here I'm just writing the file to disk, but in the real world
// would have used e.g. S3 for something like this. Again, this would have required
// non-standard packages

func ChangeRoom(robot types.Robot) error {
	room, err := GetRoom(robot.RoomID)
	if err != nil {
		return fmt.Errorf("error in ChangeRoom: %w", err)
	}

	room.XPosition = robot.XPosition
	room.YPosition = robot.YPosition
	room.Direction = robot.Direction

	if err := SaveRoom(room); err != nil {
		return fmt.Errorf("error in ChangeRoom: %w", err)
	}

	return nil
}

// GetRoom validates that the starting position of the robot is not out of bounds
func GetRoom(roomID int) (types.Room, error) {
	fileName := fmt.Sprintf("%v%v.json", dbFiles, roomID)
	dat, err := os.ReadFile(fileName)
	if err != nil {
		return types.Room{}, fmt.Errorf("error in GetRoom, cannot read file %v: %w", fileName, err)
	}

	var room types.Room
	if err := json.Unmarshal(dat, &room); err != nil {
		return types.Room{}, fmt.Errorf("error in GetRoom, cannot unmarshal data in file %v: %w", fileName, err)
	}

	return room, nil
}

// SaveRoom creates or updates a JSON file for a room with options position and direction of the robot
func SaveRoom(room types.Room) error {
	roomJSON, err := json.Marshal(room)
	if err != nil {
		return fmt.Errorf("error in SaveRoom marshaling %+v: %w", room, err)
	}

	fileName := fmt.Sprintf("%v%v.json", dbFiles, room.ID)
	file, errs := os.Create(fileName)
	if errs != nil {
		return fmt.Errorf("error in SaveRoom, could not create file for %+v: %w", room, err)
	}
	defer file.Close() // no-lint:all

	_, errs = file.WriteString(string(roomJSON))
	if errs != nil {
		return fmt.Errorf("error in SaveRoom, writing %+v to file: %w", room, err)
	}

	return nil
}
