package types

type Room struct {
	ID        int
	Width     int
	Height    int
	XPosition int
	YPosition int
	Direction string
}

type Robot struct {
	RoomID    int    `json:"roomId"`
	XPosition int    `json:"xPosition"`
	YPosition int    `json:"yPosition"`
	Direction string `json:"direction"`
}
