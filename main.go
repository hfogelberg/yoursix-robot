package main

import (
	"net/http"

	"github.com/hfogelberg/yourSixRobot/apis"
)

func main() {
	http.HandleFunc("/room", apis.RoomHandler)
	http.HandleFunc("/robot/start", apis.StartRobotHandler)
	http.HandleFunc("/robot/move", apis.MoveRobotHandler)

	http.ListenAndServe(":3000", nil) // nolint:all
}
