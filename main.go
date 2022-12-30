/*
- WUC -- WordPress Update Checker
- Author: Byron Stuike
- Release Date: 2022-10-11
- go build -o /Users/byron/Documents/bitbucket/operations/bin/wpc main.go
- GOOS=linux GOARCH=amd64 go build -o /Users/byron/Documents/bitbucket/operations/bin/wpc main.go
*/

package main

import (
	I "check/workers"
)

// Launch the program and execute according to the supplied flag
func main() {
	I.Plugin()
}
