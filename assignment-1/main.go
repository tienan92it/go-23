package main

import (
	"fmt"
	"os"
	"strings"
)

const colorRed = "\033[31m"

// unpackArgs unpacks the given arguments.
//
// It takes a slice of strings as input and returns four values: first name, last name, middle name (optional), and country code.
func unpackArgs(args []string) (string, string, string, string) {
	length := len(args)
	if length == 3 {
		return args[0], args[1], "", args[length-1]
	} else {
		return args[0], args[1], strings.Join(args[2:length-1], " "), args[length-1]
	}
}

func trimDuplicateWhitespaces(input string) string {
	// Split the string into substrings using spaces as delimiters
	// and then join the substrings with a single space
	return strings.Join(strings.Fields(input), " ")
}

func reorderName(args []string) string {
	firstName, lastName, middleName, countryCode := unpackArgs(args)
	switch strings.ToUpper(countryCode) {
	case "VN":
		return trimDuplicateWhitespaces(lastName + " " + middleName + " " + firstName)
	case "US":
		return trimDuplicateWhitespaces(firstName + " " + lastName + " " + middleName)
	default:
		return trimDuplicateWhitespaces(firstName + " " + lastName + " " + middleName)
	}
}

func main() {
	args := os.Args[1:] // Get all arguments except the program name

	if len(args) < 3 {
		fmt.Println(string(colorRed), "Error: First name, last name, and country code are required.")
	} else {
		fmt.Println("Output: ", reorderName(args))
	}
}
