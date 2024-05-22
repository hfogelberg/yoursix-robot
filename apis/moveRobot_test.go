package apis

import (
	"github.com/hfogelberg/yourSixRobot/types"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestMoveRobotAlongPath(t *testing.T) {
	tests := []struct {
		description     string
		roomIn          types.Room
		xOut            int
		yOut            int
		directionOut    string
		shouldHaveError bool
	}{
		{
			description: "Test 1",
			roomIn: types.Room{
				Width:     5,
				Height:    5,
				XPosition: 2,
				YPosition: 3,
				Direction: North,
				Path:      "FRFRFRF",
			},
			xOut:            2,
			yOut:            3,
			directionOut:    West,
			shouldHaveError: false,
		},
		{
			description: "Test 2",
			roomIn: types.Room{
				Width:     5,
				Height:    5,
				XPosition: 2,
				YPosition: 2,
				Direction: East,
				Path:      "FRRFFRF",
			},
			xOut:            1,
			yOut:            3,
			directionOut:    North,
			shouldHaveError: false,
		},
		{
			description: "Test 3",
			roomIn: types.Room{
				Width:     7,
				Height:    5,
				XPosition: 5,
				YPosition: 4,
				Direction: East,
				Path:      "FFRFFRFFLF",
			},
			xOut:            5,
			yOut:            1,
			directionOut:    South,
			shouldHaveError: false,
		},
		{
			description: "Test 4",
			roomIn: types.Room{
				Width:     3,
				Height:    3,
				XPosition: 2,
				YPosition: 2,
				Direction: East,
				Path:      "FFFFFFF",
			},
			xOut:            3,
			yOut:            2,
			directionOut:    East,
			shouldHaveError: true,
		},
	}

	for _, test := range tests {
		log.Println(test.description)
		roomOut, err := moveRobotAlongPath(test.roomIn)
		assert.Equal(t, test.xOut, roomOut.XPosition)
		assert.Equal(t, test.yOut, roomOut.YPosition)
		assert.Equal(t, test.directionOut, roomOut.Direction)
		if test.shouldHaveError {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestDoMoveForwardInDirection(t *testing.T) {
	tests := []struct {
		description string
		direction   string
		xPosIn      int
		yPosIn      int
		xPosOut     int
		yPosOut     int
	}{
		{
			description: "move north",
			direction:   North,
			xPosIn:      2,
			yPosIn:      2,
			xPosOut:     2,
			yPosOut:     3,
		},
		{
			description: "move east",
			direction:   East,
			xPosIn:      2,
			yPosIn:      2,
			xPosOut:     3,
			yPosOut:     2,
		},
		{
			description: "move south",
			direction:   South,
			xPosIn:      2,
			yPosIn:      2,
			xPosOut:     2,
			yPosOut:     1,
		},
		{
			description: "move west",
			direction:   West,
			xPosIn:      2,
			yPosIn:      2,
			xPosOut:     1,
			yPosOut:     2,
		},
	}

	for _, test := range tests {
		newX, newY := doMoveForwardInDirection(test.direction, test.xPosIn, test.yPosIn)
		assert.Equal(t, test.xPosOut, newX)
		assert.Equal(t, test.yPosOut, newY)
	}
}

func TestCanMoveForwardsInDirection(t *testing.T) {
	tests := []struct {
		description string
		direction   string
		room        types.Room
		currentX    int
		currentY    int
		canMove     bool
	}{
		{
			description: "can move north",
			direction:   North,
			room: types.Room{
				Width:  3,
				Height: 3,
			},
			currentX: 2,
			currentY: 2,
			canMove:  true,
		},
		{
			description: "can not move north",
			direction:   North,
			room: types.Room{
				Width:  3,
				Height: 3,
			},
			currentX: 2,
			currentY: 3,
			canMove:  false,
		},
		{
			description: "can move west",
			direction:   West,
			room: types.Room{
				Width:  3,
				Height: 3,
			},
			currentX: 2,
			currentY: 2,
			canMove:  true,
		},
		{
			description: "can not move west",
			direction:   West,
			room: types.Room{
				Width:  3,
				Height: 3,
			},
			currentX: 1,
			currentY: 2,
			canMove:  false,
		},
		{
			description: "can move south",
			direction:   South,
			room: types.Room{
				Width:  3,
				Height: 3,
			},
			currentX: 2,
			currentY: 2,
			canMove:  true,
		},
		{
			description: "can not move south",
			direction:   South,
			room: types.Room{
				Width:  3,
				Height: 3,
			},
			currentX: 2,
			currentY: 1,
			canMove:  false,
		},
		{
			description: "can move east",
			direction:   East,
			room: types.Room{
				Width:  3,
				Height: 3,
			},
			currentX: 2,
			currentY: 2,
			canMove:  true,
		},
		{
			description: "can not move east",
			direction:   East,
			room: types.Room{
				Width:  3,
				Height: 3,
			},
			currentX: 3,
			currentY: 2,
			canMove:  false,
		},
	}

	for _, test := range tests {
		res := canMoveForwardInDirection(test.direction, test.room, test.currentX, test.currentY)
		assert.Equal(t, test.canMove, res)
	}
}

func TestComplexMovement(t *testing.T) {
	roomIn := types.Room{
		Width:     5,
		Height:    5,
		XPosition: 3,
		YPosition: 3,
		Direction: "N",
		Path:      "FRFLFLFFLF",
	}
	expectedRoomOut := types.Room{
		Width:     5,
		Height:    5,
		XPosition: 3,
		YPosition: 3,
		Direction: "S",
	}

	roomOut, err := moveRobotAlongPath(roomIn)
	assert.Equal(t, expectedRoomOut.Direction, roomOut.Direction)
	assert.NoError(t, err)
}

// TestGetDirectionToMove returns direction robot is facing after
// rotation or movement
func TestGetDirectionToMove(t *testing.T) {
	tests := []struct {
		description       string
		direction         string
		movement          string
		expectedDirection string
		exectedAction     string
	}{
		{
			description:       "facing north, moving forward",
			direction:         North,
			movement:          MoveForward,
			expectedDirection: North,
			exectedAction:     MoveForward,
		},
		{
			description:       "facing north, turning right",
			direction:         North,
			movement:          Right,
			expectedDirection: East,
			exectedAction:     "",
		},
		{
			description:       "facing north, turning left",
			direction:         North,
			movement:          Left,
			expectedDirection: West,
			exectedAction:     "",
		},
		{
			description:       "facing east, moving forward",
			direction:         East,
			movement:          MoveForward,
			expectedDirection: East,
			exectedAction:     MoveForward,
		},
		{
			description:       "facing east, turning right",
			direction:         East,
			movement:          Right,
			expectedDirection: South,
			exectedAction:     "",
		},
		{
			description:       "facing east, turning left",
			direction:         East,
			movement:          Left,
			expectedDirection: North,
			exectedAction:     "",
		},
		{
			description:       "facing south, moving forward",
			direction:         South,
			movement:          MoveForward,
			expectedDirection: South,
			exectedAction:     MoveForward,
		},
		{
			description:       "facing south, turning right",
			direction:         South,
			movement:          Right,
			expectedDirection: West,
		},
		{
			description:       "facing south, turning left",
			direction:         South,
			movement:          Left,
			expectedDirection: East,
		},
		{
			description:       "facing west, moving forward",
			direction:         West,
			movement:          MoveForward,
			expectedDirection: West,
		},
		{
			description:       "facing west, turning right",
			direction:         West,
			movement:          Right,
			expectedDirection: North,
		},
		{
			description:       "facing west, turning left",
			direction:         West,
			movement:          Left,
			expectedDirection: South,
		},
	}

	for _, test := range tests {
		log.Println(test.description)
		resDirection := getDirectionToMove(test.direction, test.movement)
		assert.Equal(t, test.expectedDirection, resDirection)
	}
}

// TestDoRightTurn returns direction robot ois facing after right turn
func TestDoRightTurn(t *testing.T) {
	tests := []struct {
		description string
		direction   string
		movement    string
		expected    string
	}{
		{
			description: "facing north, turning right",
			direction:   North,
			movement:    Right,
			expected:    East,
		},
		{
			description: "facing east, turning right",
			direction:   East,
			movement:    Right,
			expected:    South,
		},
		{
			description: "facing south, turning right",
			direction:   South,
			movement:    Right,
			expected:    West,
		},
		{
			description: "facing west, turning right",
			direction:   West,
			movement:    Right,
			expected:    North,
		},
	}

	for _, test := range tests {
		res := doRightTurn(test.direction)
		assert.Equal(t, test.expected, res)
	}
}

// TestDoLeftTurn returns direction robot is facing after left turn
func TestDoLeftTurn(t *testing.T) {
	tests := []struct {
		description string
		direction   string
		movement    string
		expected    string
	}{
		{
			description: "facing north, turning left",
			direction:   North,
			movement:    Right,
			expected:    West,
		},
		{
			description: "facing east, turning left",
			direction:   East,
			movement:    Right,
			expected:    North,
		},
		{
			description: "facing south, turning left",
			direction:   South,
			movement:    Right,
			expected:    East,
		},
		{
			description: "facing west, turning left",
			direction:   West,
			movement:    Right,
			expected:    South,
		},
	}

	for _, test := range tests {
		res := doLeftTurn(test.direction)
		assert.Equal(t, test.expected, res)
	}
}
