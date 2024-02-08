package puppy

import (
	"fmt"

	"github.com/dzivdzi/dog"
)

var LittleDog = "LittleDog"
const BigDog = "BigDog"

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}

func From11() {
	fmt.Println("Im from version 1.1.0!")
}

func From12() {
	fmt.Println("Im from version 1.2.0!")
}

func From13() {
	fmt.Println("Im from version 1.3.0!")
}
