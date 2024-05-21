package apis

import (
	"github.com/hfogelberg/yourSixRobot/types"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStartPositionIsInBounds(t *testing.T) {
	tests := []struct {
		description string
		room        types.Room
		robot       types.Robot
		expected    bool
	}{
		{
			description: "start position in bounds",
			room: types.Room{
				Width:  5,
				Height: 5,
			},
			robot: types.Robot{
				XPosition: 3,
				YPosition: 3,
			},
			expected: true,
		},
		{
			description: "x-position too high",
			room: types.Room{
				Width:  5,
				Height: 5,
			},
			robot: types.Robot{
				XPosition: 6,
				YPosition: 3,
			},
			expected: false,
		},
		{
			description: "y-position too high",
			room: types.Room{
				Width:  5,
				Height: 5,
			},
			robot: types.Robot{
				XPosition: 3,
				YPosition: 6,
			},
			expected: false,
		},
	}

	for _, test := range tests {
		res := startPositionIsInBounds(test.room, test.robot)
		assert.Equal(t, test.expected, res)
	}
}

func TestRobotInPutIsValid(t *testing.T) {
	// No need to test all cases here since there are unit tests doing that.
	// Just check that an error is returned when it should
	tests := []struct {
		description     string
		robot           types.StartRobotIn
		shouldHaveError bool
	}{
		{
			description: "start position is valid",
			robot: types.StartRobotIn{
				RoomID:    42,
				XPosition: 5,
				YPosition: 10,
				Direction: North,
			},
			shouldHaveError: false,
		},
		{
			description: "X-position not valid",
			robot: types.StartRobotIn{
				RoomID:    42,
				XPosition: -1,
				YPosition: 10,
				Direction: North,
			},
			shouldHaveError: true,
		},
		{
			description: "Y-position not valid",
			robot: types.StartRobotIn{
				RoomID:    42,
				XPosition: 5,
				YPosition: -1,
				Direction: North,
			},
			shouldHaveError: true,
		},
		{
			description: "Direction not valid",
			robot: types.StartRobotIn{
				RoomID:    42,
				XPosition: 5,
				YPosition: 10,
				Direction: "F",
			},
			shouldHaveError: true,
		},
		{
			description: "Room ID not valid",
			robot: types.StartRobotIn{
				RoomID:    -1,
				XPosition: 5,
				YPosition: 10,
				Direction: North,
			},
			shouldHaveError: true,
		},
	}

	for _, test := range tests {
		err := robotInputIsValid(test.robot)
		// Don't check the message, just that an error is returned when it should
		if test.shouldHaveError {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestRoomIDValid(t *testing.T) {
	tests := []struct {
		description string
		roomID      int
		expected    bool
	}{
		{
			description: "valid room ID",
			roomID:      42,
			expected:    true,
		},
		{
			description: "negative room ID",
			roomID:      -1,
			expected:    false,
		},
		{
			description: "room ID is zero",
			roomID:      0,
			expected:    false,
		},
	}

	for _, test := range tests {
		res := roomIdValid(test.roomID)
		assert.Equal(t, test.expected, res)
	}
}

func TestNewRoomIsValid(t *testing.T) {
	// No need to test all cases here since there are unit tests doing that.
	// Just check that an error is returned when it should
	tests := []struct {
		description     string
		room            types.RoomSetup
		shouldHaveError bool
	}{
		{
			description: "room is valid",
			room: types.RoomSetup{
				Width:  10,
				Height: 5,
			},
			shouldHaveError: false,
		},
		{
			description: "wrong width",
			room: types.RoomSetup{
				Width:  -1,
				Height: 5,
			},
			shouldHaveError: true,
		},
		{
			description: "wrong height",
			room: types.RoomSetup{
				Width:  10,
				Height: -2,
			},
			shouldHaveError: true,
		},
	}

	for _, test := range tests {
		err := newRoomIsValid(test.room)
		// Don't check the message, just that an error is returned when it should
		if test.shouldHaveError {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestStartCooOrdinateIsValid(t *testing.T) {
	tests := []struct {
		description string
		coordinate  int
		expected    bool
	}{
		{
			description: "valid coordinate",
			coordinate:  5,
			expected:    true,
		},
		{
			description: "zero coordinate",
			coordinate:  0,
			expected:    false,
		},
		{
			description: "negative coordinate",
			coordinate:  -1,
			expected:    false,
		},
		{
			description: "very high coordinate",
			coordinate:  math.MaxUint32,
			expected:    true,
		},
	}

	for _, test := range tests {
		res := coordinateIsValid(test.coordinate)
		assert.Equal(t, test.expected, res)
	}
}

func TestDirectionIsValid(t *testing.T) {
	tests := []struct {
		direction string
		expected  bool
	}{
		{
			direction: North,
			expected:  true,
		},
		{
			direction: West,
			expected:  true,
		},
		{
			direction: South,
			expected:  true,
		},
		{
			direction: East,
			expected:  true,
		},
		{
			direction: "B",
			expected:  false,
		},
	}

	for _, test := range tests {
		res := directionIsValid(test.direction)
		assert.Equal(t, test.expected, res)
	}
}
