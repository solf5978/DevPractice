package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

const (
	Port = ":8080"
)

func serveDynamic(w http.ResponseWriter, r *http.Request) {
	response := "The Time Is Now" + time.Now().String()
	fmt.Fprintln(w, response)
}

func serveStatic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static.html")
}

// Package Level Declaration, MUST Using Full declaration
var (
	actorName string = "ElISABETH sLADEN"
	companion string = "sarah Jane Smith"
	doctorNum int    = 3
	seasion   int    = 11
)

var (
	counter = 0
)

func main() {
	// Variable Declaration
	/*
		var i int
		i = 42
		var j int = 42
		k := 42
	*/
	i := 42
	fmt.Println(i)
	fmt.Println(strconv.Itoa(seasion))

	http.HandleFunc("/static", serveStatic)
	http.HandleFunc("/", serveDynamic)
	http.ListenAndServe(Port, nil)
}
