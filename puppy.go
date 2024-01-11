package puppy

import (
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
