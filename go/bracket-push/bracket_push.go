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
			if len(stack) == 0 {
				return false, nil
			}

			end := len(stack) - 1
			p := stack[end]
			stack = stack[0:end]

			if r != p {
				return false, nil
			}
		}
	}
	return len(stack) == 0, nil
}
