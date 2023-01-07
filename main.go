/*
- WUC -- WordPress Update Checker
- Author: Byron Stuike
- Release Date: 2023-01-06
- go build -o /Users/byron/Documents/git/checker/bin/mwpc main.go
- GOOS=linux GOARCH=amd64 go build -o /Users/byron/Documents/git/checker/bin/lwpc main.go
*/

package main

import (
	I "check/workers"
)

// Launch the program and execute according to the supplied flag
func main() {
	I.Plugin()
}
