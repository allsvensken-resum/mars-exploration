# Mars Planet Exploration 


We are in a mission to control a rover that explore Mars planet. And the mission control already gave us a rover instructions to move into specific directions and positions.

## Rules

- F for moving forward 1 block.
- B for moving backward 1 block.
- L for turning left.
- R for turning right.
- HL for half turning left.
- HR for half turning right.
- A rover cannot be in negative position.
- Rover always start in position 0,0 with facing north.
- If rover reaching the edge it must maintain the direction and position.
- The list of directions that rover can have are listed below.
  - N
  - S
  - E
  - W
  - NW
  - NE
  - SW
  - SE
- The rover moving changes both x and y coordinates with value one when its direction is halfturn (NE, NW, SE, SW)

## Input

- This application is using RESTAPI as interface for this application.

- This application received file as input **ONLY ACCEPT .txt FILE**.

- The first line will always be a size of a maps represent in integer only without any seperator.

- **ONLY SUPPORT MAXIMUM MAP SIZE 99** for now.

- The other line are the list of instructions which must be one instruction per line.

- The list of instructions that you can use are listed above **Rules**.

## How to use application?

- You can easily run this application with docker-compose up
- This application is running on port **8080**.
- This feature is serving on **localhost:8080/mars/explore**, It receives file as input with **file key**.
- You can test it with postman as client to upload the file like picture below.
![Postman Example](/assets/mars-exploration-test-01.png)
- The output is **JSON** with instructions array that you write down on file and the result array of the rover direction and position. 
- The example instructions file are included in [Example files](https://github.com/allsvensken-resum/mars-exploration/tree/main/assets/examples).

### Example

  **Input**    

    24
    R
    F
    L
    F
    L
    L
    F
    R

 **Output**

    "instructions": [
        "24",
        "R",
        "F",
        "L",
        "F",
        "L",
        "L",
        "F",
        "R"
    ],
    "result": [
        "N:0,0",
        "E:0,0",
        "E:1,0",
        "N:1,0",
        "N:1,1",
        "W:1,1",
        "S:1,1",
        "S:1,0",
        "W:1,0"
    ]
