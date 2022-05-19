package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	number1, number2, err := extractTwoInt(input)
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	sum := number1 + number2
	output = strconv.Itoa(sum)

	return
}

func extractTwoInt(input string) (firstInt int, secondInt int, err error) {
	input = strings.Join(strings.Fields(input), "")

	if len(input) < 1 {
		return firstInt, secondInt, errorEmptyInput
	}

	operandsAsStrings := extractTwoAssumedOperands(input)

	if len(operandsAsStrings) != 2 {
		return firstInt, secondInt, errorNotTwoOperands
	}

	integers, err := parseIntegersFromStrings(operandsAsStrings)

	if err != nil {
		return firstInt, secondInt, err
	}

	firstInt = integers[0]
	secondInt = integers[1]

	return
}

func extractTwoAssumedOperands(input string) (operands []string) {
	operands = strings.Split(input, "+")
	if len(operands) != 2 {
		operands = strings.Split(input, "-")

		// case "X-Y"
		if len(operands) == 2 {
			operands = []string{operands[0], "-" + operands[1]}
		}

		// case "-X-Y"
		if len(operands) == 3 {
			operands = []string{"-" + operands[1], "-" + operands[2]}
		}
	}

	return
}

func parseIntegersFromStrings(input []string) (integers []int, err error) {
	for _, chunk := range input {
		foundInt, err := strconv.Atoi(chunk)
		if err != nil {
			return integers, err
		}

		integers = append(integers, foundInt)
	}

	return
}
