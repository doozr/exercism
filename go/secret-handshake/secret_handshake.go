// Package secret is all about secret handshakes
package secret

var testVersion = 1

var actions = []string{
	"wink",
	"double blink",
	"close your eyes",
	"jump",
}

// Handshake produces a secret handshake depending on the provided code
func Handshake(code uint) []string {
	var order = []uint{0, 1, 2, 3}
	handshake := make([]string, 0, 0)

	if code&(1<<4) > 0 {
		order = []uint{3, 2, 1, 0}
	}

	for _, x := range order {
		if code&(1<<x) > 0 {
			handshake = append(handshake, actions[x])
		}
	}
	return handshake
}
