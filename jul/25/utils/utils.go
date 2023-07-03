package utils

import (
	"math"
	"strconv"
)

// expreToNumber convert an expression that describes a number
// according to the rule of the problem
// Rules:
//  1. Digits are interepreted from right to left
//  2. A digit is represented as in base 5, for example 20 = 0*5^0 + 2*5^1 = 10
//     so 20 really is 10 in base 10, our normal base system
//  3. '-' represents -1 and '=' represents -2
func FromSNAFU(expr string) int {
	value := 0
	for i := len(expr) - 1; i > -1; i-- {
		exponent := len(expr) - i - 1
		power := int(math.Pow(5, float64(exponent)))
		if expr[i] == '-' {
			value -= power
		} else if expr[i] == '=' {
			value -= 2 * power
		} else {
			digit := int(expr[i] - '0')
			value += digit * power
		}
	}

	return value
}

// example: 4890 is '2=-1=0'
// division by 5 leaves 0, 1, 2, 3 and 4 as remaining
// 0 => 0
// 1 => 1
// 2 => 2
// 3 => '1='
// 4 => 1-

// 10 / 5 = 2 10%5 = 0
func ToSNAFU(number int) string {
	if number == 0 || number == 1 || number == 2 {
		return strconv.Itoa(number)
	}

	if number == -1 {
		return "-"
	}

	if number == -2 {
		return "="
	}

	d := number / 5
	r := number % 5
	if r == 3 {
		r = -2
		d++
	}
	if r == 4 {
		r = -1
		d++
	}

	return ToSNAFU(d) + ToSNAFU(r)
}
