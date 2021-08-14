package main

import (
	"log"
	"os"
)

func main() {
	key := "DB_CONN"
	//	Set the environmental variable
	os.Setenv(key, "postgres://as:as@example.com/pg?sslmod=verify-full")
	val := GetEnvDefault(key, "postgres://as:as@localhost/pg?sslmode=verify-full")
	log.Println("The value is :" + val)
}

func GetEnvDefault(key, defVal string) string {
	//	The LookupEnv function provides two values as a result
	//	the value of the variable, and the boolean value
	//	if the variable was set or not in the environment

	//	Disvantage using os.Getenv function could returns an empty string
	//	os.LookupEnv returns the string as a value of the environment variable
	//	and the boolean value that indicates whether the variable was set or not
	val, ex := os.LookupEnv(key)
	if !ex {
		return defVal
	}
	return val
}
