package workers

import (
	"fmt"
	"log"
	"os"
)

var short = []string{tmp, grp, web}

// Plugin triggers the related functions
func Plugin() {
	if contains() {
		Logging("Plugin update triggered on " + os.Args[4])
		ups := wpcli("plugin", "list", "--update=available")
		body := freebies(ups) + assemble()
		if len(body) > 0 {
			err := os.WriteFile("updates/updates.txt", []byte(body), 0666)
			trouble(err)
			mailman(body)
		} else {
			Logging("No updates found")
		}
		for _, v := range short {
			cleanup(v)
		}
	}
}

// Theme triggers the related functions
func Theme() {
	Logging("Themes update triggered on " + os.Args[4])
	ups := wpcli("theme", "list", "--update=available")
	body := freebies(ups) + assemble()
	fmt.Print(body)
}

// Core triggers the related functions
func Core() {
	Logging("Core update triggered on " + os.Args[4])
	ups := wpcli("wp", "core", "check-update")
	body := freebies(ups)
	fmt.Print(body)
}

// Logging records a message to the log file
func Logging(message string) {
	file, err := os.OpenFile("logs/wpc.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	trouble(err)
	log.SetOutput(file)
	log.Println(message)
}
