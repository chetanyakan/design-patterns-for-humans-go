// It’s impossible to implement the classic Factory Method pattern in Go
// due to the lack of OOP features such as classes and inheritance.
// However, we can still implement the basic version of the pattern, the Simple Factory.

package main

import (
	"fmt"
)

type Door interface {
	getWidth() float64
	getHeight() float64
}

type WoodenDoor struct {
	width  float64
	height float64
}

func (w *WoodenDoor) getWidth() float64 {
	return w.width
}

func (w *WoodenDoor) getHeight() float64 {
	return w.height
}

type DoorFactory struct{}

func (d *DoorFactory) makeWoodenDoor(width, height float64) WoodenDoor {
	return WoodenDoor{width, height}
}

func main() {
	doorFactory := &DoorFactory{}

	// Make me a door of 100x200
	door := doorFactory.makeWoodenDoor(100, 200)

	fmt.Println("Width: ", door.getWidth())
	fmt.Println("Height: ", door.getHeight())

	// Make me a door of 50x100
	//door2 := doorFactory.makeWoodenDoor(50, 100)
}
