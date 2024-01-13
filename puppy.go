package puppy

import (
	"fmt"

	"github.com/YutaUtah/dog"
)

// go mod init github.com/YutaUtah/puppy
func Bark() string {
	return "woof!!"
}

func Barks() string {
	return "woof!woof!woof!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}

func From11() {
	fmt.Println("I'm from version 1.1.0")
}

func From12() {
	fmt.Println("I'm from version 1.2.0")
}

func From13() {
	fmt.Println("I'm from version 1.3.0")
}
