package main

import "fmt"

type Door struct{
	state bool
}
const OPEN = true
const CLOSE =false

func OpenDoor(ptrDoor *Door){
	fmt.Println("Door Openning...")
	ptrDoor.state = OPEN
}

func CloseDoor(ptrDoor *Door){
	fmt.Println("Door Closing...")
	ptrDoor.state = CLOSE
}

func IsDoorOpen(ptrDoor Door) bool {
	fmt.Println("is the Door opened ?")
	return ptrDoor.state
}

func IsDoorClose(ptrDoor Door) bool {
	fmt.Println("is the Door closed ?")
	return ptrDoor.state
}

func main() {
	door := Door{}

	OpenDoor(&door)
	if IsDoorClose(door) {
		OpenDoor(&door)
	}
	if IsDoorOpen(door) {
		CloseDoor(&door)
	}
	if door.state == OPEN {
		CloseDoor(&door)
	}
}
