package utils

import (
	"fmt"
	"log"
)

func ErrorPanic(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func ErrorLog(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ErrorPrint(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
