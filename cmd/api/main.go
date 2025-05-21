package main

import (
	"fmt"
	"log"
)

// this is a version
const version = "1.0.0"

// this is config for the app ,like port and eviromet
type Config struct {
	port int
	env  string
}

// This is a struct that holds dependencies for the application like  database, for now only config and logger
type Appilcation struct {
	config Config
	logger *log.Logger
}

func main() {

	fmt.Println("Hello world")
}
