package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	// Read input from standard input
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := scanner.Text()

		// Call the ReverseString function
		output := ReverseString(input)

		// Print the result
		fmt.Println(output)
	}
}

// ReverseString returns the reversed string of s.
func ReverseString(s string) string {
	result := strings.Builder{}

	for i, w := len(s), 0; i > 0; i -= w {
		r, width := utf8.DecodeLastRuneInString(s[:i])
		w = width

		result.WriteRune(r)
	}

	return result.String()
}
