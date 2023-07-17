package puppy

import (
	"github.com/shubham-baliyan/go-dog"
)

func Bark() string {
	return "Woof!"
}
func Barks() string {
	return "Woof Woof Woof !!!"
}

func BigBark() string {
	return dog.WhenGrownUp(Barks())
}
