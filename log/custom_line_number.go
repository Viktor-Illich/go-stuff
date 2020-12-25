package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	// LstdFlags = Ldate | Ltime
	// initial values for the standard logger
	iLog := log.New(f, "customLogLineNumber ", log.LstdFlags)

	// Print the line number of when and where something has happened
	iLog.SetFlags(log.LstdFlags | log.Lshortfile)
	iLog.Println("Hello there!")
	iLog.Println("Another log entry!")
}
