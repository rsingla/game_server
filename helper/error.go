package helper

import (
	"log"
)

// Error is a helper function to log errors
func Error(err error) {
	if err != nil {
		log.Println(err)
	}
}

// ErrorFatal is a helper function to log fatal errors
func ErrorFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// ErrorPanic is a helper function to log panic errors
func ErrorPanic(err error) {
	if err != nil {
		log.Panic(err)
		panic(err)
	}
}
