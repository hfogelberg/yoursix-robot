package apis

import (
	"fmt"
	"github.com/hfogelberg/yourSixRobot/types"
)

const (
	North = "N"
	East  = "E"
	South = "S"
	West  = "W"
)

func startPositionIsInBounds(room types.Room, robot types.Robot) bool {
	if robot.XPosition > room.Width {
		return false
	}

	if robot.YPosition > room.Height {
		return false
	}

	return true
}

func robotInputIsValid(robot types.StartRobotIn) error {
	ok := coordinateIsValid(robot.XPosition)
	if !ok {
		return fmt.Errorf("info: xPosition %v in robotInputIsValid is not valid", robot.XPosition)
	}

	ok = coordinateIsValid(robot.YPosition)
	if !ok {
		return fmt.Errorf("info: yPosition %v in robotInputIsValid is not valid", robot.YPosition)
	}

	ok = directionIsValid(robot.Direction)
	if !ok {
		return fmt.Errorf("info: direction %v in robotInputIsValid is not valid", robot.Direction)
	}

	ok = roomIdValid(robot.RoomID)
	if !ok {
		return fmt.Errorf("info: room ID %v in robotInputIsValid is not valid", robot.RoomID)
	}

	return nil
}

func roomIdValid(roomID int) bool {
	if roomID < 1 {
		return false
	}

	return true
}

func newRoomIsValid(setup types.RoomSetup) error {
	ok := coordinateIsValid(setup.Height)
	if !ok {
		return fmt.Errorf("info: height in newRoomIsValid %v is not valid for new room", setup.Height)
	}

	ok = coordinateIsValid(setup.Width)
	if !ok {
		return fmt.Errorf("info: width in newRoomIsValid %v is not valid for new room", setup.Height)
	}

	return nil
}

func coordinateIsValid(coordinate int) bool {
	// The grid in the application is 1-based.
	// The maximum grid size is only limited by the maximum value of a Go integer
	if coordinate < 1 {
		return false
	}

	return true
}

func directionIsValid(direction string) bool {
	switch direction {
	case North:
		return true
	case East:
		return true
	case South:
		return true
	case West:
		return true
	default:
		return false
	}
}
