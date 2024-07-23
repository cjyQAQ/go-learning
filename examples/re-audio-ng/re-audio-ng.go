package reaudiong

import (
	"fmt" 
	"examples/ae"
)


func Start() {
	fmt.Println("ng start")
}


func Stop() {
	fmt.Println("ng stop")
}


func DealData() {
	InsertIntoMysql()
	ae.InsertIntoRedis()
}