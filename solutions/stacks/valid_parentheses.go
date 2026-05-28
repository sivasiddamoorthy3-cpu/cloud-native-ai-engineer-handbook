
package stacks

func IsValid(s string) bool {
    stack := []rune{}

    pairs := map[rune]rune{
        ')':'(',
        ']':'[',
        '}':'{',
    }

    for _, ch := range s {
        switch ch {
        case '(', '[', '{':
            stack = append(stack, ch)
        default:
            if len(stack) == 0 {
                return false
            }

            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]

            if top != pairs[ch] {
                return false
            }
        }
    }

    return len(stack) == 0
}
