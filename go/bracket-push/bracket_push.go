package brackets

var testVersion = 4

func Bracket(s string) (bool, error) {
	stack := make([]rune, 0)

	push := func(r rune) {
		stack = append(stack, r)
	}

	pop := func() (r rune) {
		if len(stack) == 0 {
			return
		}
		end := len(stack) - 1
		r = stack[end]
		stack = stack[0:end]
		return
	}

	for _, r := range s {
		switch r {
		case '{':
			push('}')
		case '[':
			push(']')
		case '(':
			push(')')
		default:
			if r != pop() {
				return false, nil
			}
		}
	}
	return len(stack) == 0, nil
}
