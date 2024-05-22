package apis

import (
	"fmt"
	"github.com/hfogelberg/yourSixRobot/types"
)

func moveRobotAlongPath(room types.Room) (types.Room, error) {
	direction := room.Direction
	currentX := room.XPosition
	currentY := room.YPosition

	for _, m := range room.Path {
		movement := string(m)

		if movement == MoveForward {
			ok := canMoveForwardInDirection(direction, room, currentX, currentY)
			if !ok {
				return setRoomOut(room, direction, currentX, currentY), fmt.Errorf("cannot move %v in romm %+v", direction, room)
			}
			currentX, currentY = doMoveForwardInDirection(direction, currentX, currentY)
		} else {
			direction = getDirectionToMove(direction, movement)
		}
	}

	room = setRoomOut(room, direction, currentX, currentY)

	return room, nil
}

func setRoomOut(room types.Room, direction string, currentX int, currentY int) types.Room {
	room.Path = ""
	room.XPosition = currentX
	room.YPosition = currentY
	room.Direction = direction

	return room
}

func doMoveForwardInDirection(direction string, xPos int, yPos int) (int, int) {
	switch direction {
	case North:
		yPos += 1
	case East:
		xPos += 1
	case South:
		yPos -= 1
	case West:
		xPos -= 1
	}
	return xPos, yPos
}

func canMoveForwardInDirection(direction string, room types.Room, currentX int, currentY int) bool {
	switch direction {
	case North:
		if currentY == room.Height {
			return false
		}
	case East:
		if currentX == room.Width {
			return false
		}
	case South:
		if currentY == 1 {
			return false
		}
	case West:
		if currentX == 1 {
			return false
		}
	default:
		return true
	}

	return true
}

func getDirectionToMove(direction string, movement string) (newDirection string) {
	if movement == MoveForward {
		return direction
	} else if movement == Right {
		return doRightTurn(direction)
	} else {
		return doLeftTurn(direction)
	}
}

func doRightTurn(direction string) string {
	switch direction {
	case North:
		return East
	case East:
		return South
	case South:
		return West
	default:
		return North
	}
}

func doLeftTurn(direction string) string {
	switch direction {
	case North:
		return West
	case East:
		return North
	case South:
		return East
	default:
		return South
	}
}
