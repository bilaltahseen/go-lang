package main

import (
	"fmt"
	"os"
	"strconv"

	errorhandling "github.com/bilaltahseen/go-lang/demo-1/src/ErrorHandling"
)

// func main() {
// 	username, password := os.Args[1], os.Args[2]
// 	result := conditions.Auth(username, password)

// 	fmt.Println("User authentication", result)
// }

// func main() {

// 	feet, err := strconv.ParseFloat(os.Args[1], 64)

// 	if err != nil {
// 		fmt.Println("Error", err)
// 		return
// 	}

// 	result := errorhandling.FeetToMeter(feet)

// 	fmt.Println("Result", result)
// }

func main() {
	// Using simple statement
	// err will be scoped based variable
	if feet, err := strconv.ParseFloat(os.Args[1], 64); err == nil {
		meter := errorhandling.FeetToMeter(feet)
		fmt.Println("Feet to meter is", meter)
	}
}
