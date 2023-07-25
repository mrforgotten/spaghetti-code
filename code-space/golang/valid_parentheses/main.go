package main

import "fmt"

/*
func isValid(s string) bool {

		if len(s) < 1 {
			return false
		}
		stack := make([]string, 0, len(s))
		tempS := s
		stackLen := 0
		for i := 0; i < len(s); i++ {
			stack = append(stack, string(tempS[0]))
			tempS = s[i+1:]
			stackLen = len(stack)
			if stackLen > 1 {
				isPair := false
				switch stack[stackLen-1] {
				case "}":
					if stack[stackLen-2] == "{" {
						isPair = true
					} else {
						return false
					}
				case ")":
					if stack[stackLen-2] == "(" {
						isPair = true
					} else {
						return false
					}
				case "]":
					if stack[stackLen-2] == "[" {
						isPair = true
					} else {
						return false
					}
				default:
					continue
				}

				if isPair {
					stack = stack[:stackLen-2]
				}
			}
		}

		return len(stack) < 1
	}
*/
var parenthesesMap = map[byte]byte{
	'}': '{',
	')': '(',
	']': '[',
}

func isValid(s string) bool {
	sBytes := []byte(s)
	stack := make([]byte, 0, len(s))
	stackLen := 0

	for _, v := range sBytes {
		stack = append(stack, v)
		stackLen = len(stack)
		if stackLen <= 1 {
			continue
		}
		v, isExist := parenthesesMap[stack[stackLen-1]]
		if isExist && v == stack[stackLen-2] {
			stack = stack[:stackLen-2]
		}
	}

	return len(stack) == 0
}

func main() {
	// s := "()[]{}"
	// s := "]"
	s := "()"

	fmt.Println(isValid(s))

}
