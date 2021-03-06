package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type ArrayValue []string

//	Require implement Custom Type
//	flag.Value interfce to be able to
//	use it in flag.Var function
func (s *ArrayValue) String() string {
	return fmt.Sprintf("%v", *s)
}

func (s *ArrayValue) Set(s string) error {
	*a = strings.Split(s, ",")
	return nil
}

func main() {
	//  Extracting flag values w/ methods returning pointers
	retry := flag.Int("retry", -1, "Defines max retry count")

	//	Read the flag using XXXVar function.
	//	Variable has to be defined prior to the flag

	var logPrefix string
	flag.StringVar(&logPrefix, "prefix", "", "Logger prefix")

	var arr ArrayValue
	flag.Var(&arr, "array", "input array to iterate through.")

	//  Execute the flag.Parse to read the flags to defined variables.
	//  without this call the flag while variables remain empty
	flag.Parse()

	//  Sample logic not related to flags
	logger := log.New(os.Stdout, logPrefix, log.Ldate)

	retryCount := 0
	for retryCount < *retry {
		logger.Println("Retrying connection")
		logger.Printf("Sending array %v\n", arr)
		retryCount++
	}
}
