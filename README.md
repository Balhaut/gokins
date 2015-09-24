# gokins
Experimenting with Bitbucket-Jenkins hook utility to learn the Go language.

# usage
`go get github.com/gorilla`
`go build gokins.go`
`./gokins`

It will be listening on port 9001, and expecting a bitbucket 2.0 webhook request. 
Right now it just has a simple struct to unmarshal the json to, mapping the basic fields in the body: repo name and branch name.

