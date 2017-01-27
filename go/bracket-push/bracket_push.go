package brackets

var testVersion = 4

// Bracket checks that all brackets are matched
func Bracket(s string) (bool, error) {
	stack := make([]rune, 0)

	for _, r := range s {
		switch r {
		case '{':
			stack = append(stack, '}')
		case '[':
			stack = append(stack, ']')
		case '(':
			stack = append(stack, ')')
		default:
			end := len(stack) - 1
			if end == -1 || r != stack[end] {
				return false, nil
			}

			stack = stack[0:end]
		}
	}
	return len(stack) == 0, nil
}
