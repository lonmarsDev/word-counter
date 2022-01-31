# Word counter server



## Features

 - Accepts as input a body of text
 - Returns the top ten most-used words along with how many times they occur in the text.
 

## Tech

Dillinger uses a number of open source projects to work properly:

- [Golang] 

## Installation

Install the dependencies and devDependencies and start the server.

```sh
go mod tidy
go run main.go
```
Testing:

Verify the deployment by navigating to localhost:8081
your preferred browser.

or via CURL
```sh
curl -X POST \
  http://localhost:8081/ \
  -H 'cache-control: no-cache' \
  -H 'postman-token: 445703cc-6dc7-a949-71f2-661f2febabfc' \
  -d 'mine break mine testing testing testing test test we able run bike school school when we are go to school, into the holy testing'
```
