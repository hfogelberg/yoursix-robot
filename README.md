# Yoursix Robot Assignment

## Description
This code is a programming assignment for Your six, running a web server which creates a grid, positions a robot in the grid, moves it and reports the position. The solution only contains backend code, written in Go.

## Testing and running
- To run tests: `make test`
- To start web server on port 3000: `make run`.

## APIS
### Room
A new grid-based room is created. Note that the room is one-based with the lowest numbers to the left and bottom (i.e. furthest west and south are one).

The endpoint receives a POST request to create a new room and saves the room and initial positions in the file-based "database".

#### Request
The endpoint requires two parameters in the body; width and height.
Sample: 
```curl
curl -i -X POST "http://localhost:3000/room" \
-d '{
    "width": 7,
    "height": 5
}'
```

#### Response
The API will on success respond with a JSON object containing the room id, used in subsequent requests. Sample: 
```json
{
  "roomId": 3508300940
}
```
Possible error codes:
- 405 Method not allowed: Other HTTP methods than POST or OPTIONS was used
- 400 Bad request: The body was malformed in some way
- 406 Not acceptable: Mandatory params missing from the request body
- 500 Internal server error: A technical issue occurred


### Start Robot
The endpoint takes positions the robot in the grid.

#### Request
The endpoint takes four mandatory parameters:
- roomId: integer of room, returned by room handler
- xPosition: x-position of robot
- yPosition: y-position of robot
- direction: direction robot is facing, permitted values are N, E, S or W.

Sample:
```curl
curl -i -X PUT http://localhost:3000/robot/start \
-d '{
	"roomId": 3508300940,
	"xPosition": 5,
	"yPosition": 4,
	"direction": "E"
}'
```
#### Response
The API will respond with a status code only
- 200 OK: Robot was placed successfully in the room
- 405 Method not allowed: Other HTTP methods than POST or OPTIONS was used
- 400 Bad request: The body was malformed in some way
- 406 Not acceptable: One or more of the body params are missing or invalid
- 409 Conflict: The robot's starting position is outside the grid
- 500 Internal server error: A technical error has occurred

### Move Robot
Move the report by sending a path, containing direction changes or moving forward

#### Request
The API responds to a PUT request, containing room id and path to move along in the request body.

Sample:
```curl
curl -i -X PUT http://localhost:3000/robot/move \
-d '{
    "roomId": 3508300940,
    "path": "FFRFFRFFLF"
}'
```

#### Response 
On success the API will respond with a JSON body, containing the new position of the robot. The first param is the X-position, the second the y-position and finally the direction the robot is facing.

Sample:
```json
{
  "report":"Report: 5 1 S"
}
```

On error one of the following status codes is returned:
- 405 Method not allowed: Other HTTP methods than POST or OPTIONS was used
- 400 Bad request: The body was malformed in some way
- 421 Misdirected request: The robot has reached the edge of the room and cannot fully follow the path
- 500 Internal server error: A technical error has occurred


