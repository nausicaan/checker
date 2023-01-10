/*
- Author: Byron Stuike
- Release Date: 2023-01-06
- go build -o /{path_to_file}/wpp-check main.go
- GOOS=linux GOARCH=amd64 go build -o /{path_to_file}/wpp-check main.go
*/

package main

import (
	W "check/workers"
)

// Launch the program and execute according to the supplied flag
func main() {
	W.Plugin()
}
