package algo

import "fmt"

func Linter(str string) error {

	s := Stack{}
	lintMapOpen := map[string]bool{
		"{": false,
		"(": false,
		"[": false,
	}

	lintMapClose := map[string]string{
		"}": "{",
		")": "(",
		"]": "[",
	}

	for _, c := range str {
		letter := string(c)
		_, ok := lintMapOpen[letter]
		if ok {
			s.Push(letter)
			continue
		}

		openBrace, ok := lintMapClose[letter]
		if ok {
			pop, err := s.Pop()
			if err != nil {
				return err
			}
			if openBrace != pop {
				return fmt.Errorf("mismatching braces")
			}
		}
	}

	if len(s) > 0 {
		return fmt.Errorf("missing closing brace")
	}

	return nil
}

type Stack []string

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() (string, error) {

	if len(*s) == 0 {
		return "", fmt.Errorf("no values in stack.  cannot pop until a value is added to stack")
	}

	result := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return result, nil
}

func Reverser(str string) string {

	s := Stack{}

	for _, c := range str {
		letter := string(c)
		s.Push(letter)
	}

	value := ""

	for i := 0; i < len(str); i++ {
		pop, err := s.Pop()

		if err != nil {
			fmt.Println(err)
		}

		value += pop
	}
	return value
}
