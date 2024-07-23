package main


import (
	"fmt"
	"examples/reaudiong"
)

func main() {
	fmt.Println("re-audio start")
	reaudiong.Start()
	reaudiong.DealData()
	reaudiong.Stop()
}