package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// String functions

func Ord(x byte) byte {
	return x
}

func Chr(x byte) string {
	return string([]byte{x})
}

func Length(s string) int16 {
	return int16(len(s))
}

func UpCase(b byte) byte {
	if b >= 'a' && b <= 'z' {
		return b - ('a' - 'A')
	}
	return b
}

func UpCaseString(input string) string {
	b := make([]byte, len(input))
	for i := 0; i < len(input); i++ {
		b[i] = UpCase(input[i])
	}
	return string(b)
}

func Copy(s string, index, count int16) string {
	if index < 1 {
		index = 1
	}
	if count > int16(len(s))-index+1 {
		count = int16(len(s)) - index + 1
	}
	return s[index-1 : index-1+count]
}

func Pos(b byte, s string) int16 {
	return int16(strings.IndexByte(s, b) + 1)
}

func Val(s string, code *int16) int16 {
	// Skip leading spaces
	orig := s
	for len(s) > 0 && s[0] == ' ' {
		s = s[1:]
	}

	// Handle '-' or '+' sign
	negative := false
	if len(s) > 0 {
		switch s[0] {
		case '-':
			negative = true
			s = s[1:]
		case '+':
			s = s[1:]
		}
	}

	// Convert decimal digits
	val := int16(0)
	gotDigits := false
	for len(s) > 0 && s[0] >= '0' && s[0] <= '9' {
		val = val*10 + int16(s[0]-'0')
		gotDigits = true
		s = s[1:]
	}

	// Error if we didn't get any digits or there are chars left
	if !gotDigits || len(s) > 0 {
		*code = int16(len(orig) - len(s) + 1)
		return 0
	}

	if negative {
		val = -val
	}
	*code = 0 // Code of zero means no error
	return val
}

func Str(n int16) string {
	return strconv.Itoa(int(n))
}

func StrWidth(n, width int16) string {
	return fmt.Sprintf("%*d", width, n)
}

func Delete(s string, index, count int16) string {
	return s[:index-1] + s[index-1+count:]
}

// Replace byte at 1-based index with b and return new string
func Replace(s string, index int16, b byte) string {
	return s[:index-1] + string([]byte{b}) + s[index:]
}

// Misc functions

var Time int16 // TODO

func Delay(milliseconds int16) {
	time.Sleep(time.Duration(milliseconds) * time.Millisecond)
}

func Sound(x uint16) {
	// TODO
}

func NoSound() {
	// TODO
}

// Math functions

func Random(end int16) int16 {
	return int16(rand.Intn(int(end)))
}

func Sqr(n int16) int16 {
	return n * n
}

func Ln(x float64) float64 {
	return math.Log(x)
}

func Exp(x float64) float64 {
	return math.Exp(x)
}

func Trunc(x float64) int16 {
	return int16(x)
}

func BoolToInt(b bool) int16 {
	if b {
		return 1
	}
	return 0
}
