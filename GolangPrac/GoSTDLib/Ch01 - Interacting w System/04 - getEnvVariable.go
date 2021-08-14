package main

import (
	"log"
	"os"
)

func main() {
	connStr := os.Getenv("windir")
	log.Printf("Connection String: %s\n", connStr)
}
