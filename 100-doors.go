package main

import (
	"fmt"
)

var (
	numberOfDoors = 10
	doors = make([]bool,10)
)
func flipDoor(currentDoor int) {
	for j := currentDoor; j < numberOfDoors; j += currentDoor + 1 {
		doors[j] = !doors[j]
	}
}

func main() {
	
	for currentDoor := 0; currentDoor < numberOfDoors; currentDoor++ {
		flipDoor(currentDoor)
	}
	fmt.Println(doors)

}