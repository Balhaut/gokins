# gokins
Experimenting with Bitbucket-Jenkins hook utility to learn the Go language and play with some web app frameworks.

# usage
`go get github.com/gorilla`
`go build gokins.go`
`./gokins`

It will be listening on port 9001, and expecting a bitbucket 2.0 webhook request. 
A simple struct exists, and the json request body is decoded to return a few basic fields for now (repo name, branch name).
