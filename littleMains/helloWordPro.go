package main

import (
	"log"
	"os"
	"os/user"
)

func main () {
	u, _ := user.Current()
	host, _ := os.Hostname()
	log.Println("Hi ", u.Username, " from ", host)	
}

