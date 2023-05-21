package util

import (
	"log"
	"os"
)

var Logger *log.Logger

func init() {
	// Clear log file
	if os.IsNotExist(os.MkdirAll("C:\\logs", 0755)) {
		panic("Unable to create logs directory")
	}
	os.Remove("C:\\logs\\debug.log")
	// Create log file
	f, err := os.OpenFile("C:\\logs\\debug.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	Logger = log.New(f, "Logger:\t", log.Ldate|log.Ltime|log.Lshortfile)
}