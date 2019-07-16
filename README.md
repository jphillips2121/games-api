# Games API 

And API for creating, retrieving, updating, and deleting data for Video Games. 

## Requirements
In order to run this API locally you will need to install the following: 

- [Go](https://golang.org/doc/install)
- [Git](https://git-scm.com/downloads)

## Running the application locally 

1. Clone this repository: `go get github.com/jphillips2121/games-api`
1. Locate the downloaded repository in `$GOPATH/src/github.com/jphillips2121/games-api`
1. Add AWS credentials to the `games_env` file. These are explained in the Configuration below 
1. run ./start.sh

## Configuration

Variable                      | Default         | Description
:-----------------------------|:----------------|:--------------------------------
`AWS_ID`                      |                 | Aws access key ID
`AWS_SECRET`                  |                 | Aws secret access key
`AWS_TOKEN`                   |                 | Aws access token (should be left blank)
`FILE_NAME`                   | developers.json | Name of File containing valid developers
`AWS_BUCKET`                  |                 | Name of bucket `FILE_NAME` is located
`AWS_REGION`                  |                 | Region of AWS bucket

## Endpoints

The API listens on port 8080, so all paths should be prefixed with `http://localhost:8080`

Method       | Path                                            | Description
:------------|:------------------------------------------------|:-------------------------------------------
**POST**     | /games                                          | Creates a new video games if valid and adds to database
**GET**      | /games                                          | Returns a list of all video games on the database
**GET**      | /games/{game_id}                                | Returns a specific video game on the database
**PUT**      | /games/{game_id}                                | Updates a video game on the database if the new data is valid
**DELETE**   | /games/{game_id}?developer={developer}          | Deletes a video game on the database if the developer is valid


