package gomodule

import (
	"fmt"

	"github.com/rssllyn/thirdlevel"
)

// Hello returns a greeting for the named person.
func Hello() string {
	// Return a greeting that embeds the name in a message.
	name := thirdlevel.Name()
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
