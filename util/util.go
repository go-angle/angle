package util

import "os"

// Hostname returns hostname
func Hostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknow"
	}
	return hostname
}
