package main

import (
	"fmt"
)

// func main() {

// 	fmt.Println(time.Now())
// 	fmt.Println("UTC time:", time.Now().UTC().Format("2006-01-02 15:04"))
// }

// package main

// import "fmt"

func main() {
	var first int = 10
	var cond int = 4

	if first <= 0 {

		fmt.Printf("first is less than or equal to 0\n")
	} else if first > 0 && first < 5 {

		fmt.Printf("first is between 0 and 5\n")
	} else {

		fmt.Printf("first is 5 or greater\n")
	}
	if cond = 5; cond > 10 {

		fmt.Printf("cond is greater than 10\n")
	} else {

		fmt.Printf("cond is not greater than 10\n")
	}

	if aoe := 4; aoe > 3 {
		fmt.Println("aoe is greater than 3")
	}
	// fmt.Println("aoe is ", aoe)
}
