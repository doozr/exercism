package wordy

import (
	"strconv"
	"strings"
)

var testVersion = 1

// parseNumber parses a number at the start of a string
func parseNumber(question string) (number int, remaining string, ok bool) {
	parts := strings.SplitN(question, " ", 2)
	number, err := strconv.Atoi(parts[0])
	if err == nil {
		ok = true
	}
	if len(parts) > 1 {
		remaining = strings.TrimLeft(parts[1], " ")
	}
	return
}

// Operator represents binary integer maths operators
type Operator func(int, int) int

var operators = map[string]Operator{
	"plus":          func(a, b int) int { return a + b },
	"minus":         func(a, b int) int { return a - b },
	"multiplied by": func(a, b int) int { return a * b },
	"divided by":    func(a, b int) int { return a / b },
}

// parseOperation parses an operator name at the start of a string and returns the equivalent function
func parseOperation(question string) (operation Operator, remaining string, ok bool) {
	for name, operator := range operators {
		if !strings.HasPrefix(question, name) {
			continue
		}

		operation = operator
		remaining = strings.TrimLeft(question[len(name):], " ")
		ok = true
		return
	}
	return
}

// Answer a wordy maths question
func Answer(question string) (answer int, ok bool) {
	if !strings.HasPrefix(question, "What is ") {
		return
	}

	question = strings.TrimRight(question[8:], " ?")

	answer, question, ok = parseNumber(question)
	var operation Operator
	var number int
	for len(question) > 0 {
		operation, question, ok = parseOperation(question)
		if !ok {
			return
		}

		number, question, ok = parseNumber(question)
		answer = operation(answer, number)
	}

	return
}
