package main

import (
	"fmt"
	"time"
)

func getOsName() string {
	return "mac"
}

func main() {
	//os := "mac"
	switch os := getOsName(); os {
	case "mac":
		fmt.Println("Mac!!")
	case "windows":
		fmt.Println("windows!!")
	default:
		fmt.Println("Default!!", os)
	}

	t := time.Now()
	fmt.Println(t.Hour())
	switch {
	case t.Hour() < 12:
		fmt.Println("Morning")
	case t.Hour() < 17:
		fmt.Println("Afternoon")
	}
}
