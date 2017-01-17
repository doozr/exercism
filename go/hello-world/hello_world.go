// Package greeting contains functions to generate greetings.
package greeting

import "fmt"

// testVersion identifies the version of the test program that you are
// writing your code to. If the test program changes in the future --
// after you have posted this code to the Exercism site -- nitpickers
// will see that your code can't necessarily be expected to pass the
// current test suite because it was written to an earlier test version.
const testVersion = 3

// HelloWorld says hello to the given name.
// If the name is empty, it says Hello to the world.
func HelloWorld(name string) (greeting string) {
	if name == "" {
		name = "World"
	}

	greeting = fmt.Sprintf("Hello, %s!", name)
	return
}
