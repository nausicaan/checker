package workers

import (
	"os"
	"os/exec"
	"strings"
)

const (
	user = "deploy"
)

// Allow only a predefined set of servers
var servers = []string{"cactuar.dmz", "coeurl.dmz", "chimera.dmz", "moogle.dmz", "moblin.dmz", "mimic.dmz"}

// Run the wp update check command
func wpcli(x, y, z string) []string {
	c := rtnByte(exec.Command("wp", x, y, z, "--fields=name,version,update_version", "--format=csv", "--ssh="+user+"@"+os.Args[2]+":"+os.Args[3], "--url="+os.Args[4]))
	f := strings.ReplaceAll(string(c), "\n", ",")
	r := strings.Split(f, ",")
	return r
}

// Format the output of free plugin updates
func freebies(r []string) string {
	var value string

	for a := 1; a < 4; a++ {
		r = append(r[:0], r[0+1:]...)
	}

	for i := 0; i < len(r)-1; i++ {
		switch r[i] {
		case "events-virtual":
			value += "bcgov-plugin/" + r[i] + ":" + r[i+2] + "\n"
		case "events-calendar-pro":
			value += "bcgov-plugin/" + r[i] + ":" + r[i+2] + "\n"
		default:
			value += "wpackagist-plugin/" + r[i] + ":" + r[i+2] + "\n"
		}
		i += 2
	}
	return strings.TrimRight(value, " ")
}
