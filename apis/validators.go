package apis

import (
	"fmt"
	"github.com/hfogelberg/yourSixRobot/types"
)

const (
	North       = "N"
	East        = "E"
	South       = "S"
	West        = "W"
	Left        = "L"
	Right       = "R"
	MoveForward = "F"
)

func moveInputIsValid(moveIn types.MoveRobotIn) error {
	if ok := roomIdValid(moveIn.RoomID); !ok {
		return fmt.Errorf("info: invalid room id %+v in moveInputIsValid", moveIn)
	}

	if ok := pathIsValid(moveIn.Path); !ok {
		return fmt.Errorf("info: invalid path id %+v in moveInputIsValid", moveIn)
	}
	return nil
}

// dummy check that path only includes valid characters, i.i. L, F and R
func pathIsValid(path string) bool {
	if len(path) == 0 {
		return false
	}

	for _, ch := range path {
		switch string(ch) {
		case Left:
		case Right:
		case MoveForward:
		default:
			return false
		}
	}

	return true
}

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
	return roomID >= 1
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
	return coordinate >= 1
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
