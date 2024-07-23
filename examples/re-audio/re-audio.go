package main


import (
	"fmt"
	"examples/re-audio-ng"
)

func main() {
	fmt.Println("re-audio start")
	reaudiong.Start()
	reaudiong.DealData()
	reaudiong.Stop()
}