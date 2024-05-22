package types

type RoomSetup struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type StartRobotIn struct {
	RoomID    int    `json:"roomId"`
	XPosition int    `json:"xPosition"`
	YPosition int    `json:"yPosition"`
	Direction string `json:"direction"`
}

type MoveRobotIn struct {
	RoomID int    `json:"roomId"`
	Path   string `json:"path"`
}
