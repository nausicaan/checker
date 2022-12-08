/*
- WUC -- WordPress Update Checker
- Author: Byron Stuike
- Release Date: 2022-10-11
- go build -o /Users/byron/Documents/bitbucket/operations/bin/wpc main.go
- GOOS=linux GOARCH=amd64 go build -o /Users/byron/Documents/bitbucket/operations/bin/wpc main.go
*/

package main

import (
	I "go/check/workers"
	"os"
)

// Launch the program and execute according to the supplied flag
func main() {
	argLength := len(os.Args)
	if argLength == 5 {
		choice := os.Args[1]
		switch choice {
		case "-p":
			I.Plugin()
		case "-t":
			I.Theme()
		case "-c":
			I.Core()
		default:
			I.Logging("Incorrect flag detected - program halted")
		}
		I.Logging("Program exited without errors")
	} else {
		I.Logging("Insufficient arguments supplied - program halted")
	}
}
