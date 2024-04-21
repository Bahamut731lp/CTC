package main

import "fmt"

func main() {
	station := NewStation("gas", 3, 5)
	fmt.Println(station.StationType)
}
