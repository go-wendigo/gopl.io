// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/ex3.11
//	$ ./comma +1.1 -12.21 +123.321 -1234.4321 +1234567890.0987654321
// 	+1.1
// 	-12.21
// 	+123.321
// 	-1,234.4321
// 	+1,234,567,890.0987654321
//
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

// !+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	var buf bytes.Buffer

	var mantissaStart int
	if s[0] == '+' || s[0] == '-' {
		buf.WriteByte(s[0])
		mantissaStart = 1
	}

	mantissaEnd := strings.Index(s, ".")
	if mantissaEnd == -1 {
		mantissaEnd = len(s)
	}
	if mantissaEnd-mantissaStart <= 3 {
		buf.WriteString(s[mantissaStart:mantissaEnd])
	} else {
		firstDigit := (mantissaEnd - mantissaStart) % 3
		if firstDigit == 0 {
			firstDigit = 3
		}
		buf.WriteString(s[mantissaStart : mantissaStart+firstDigit])
		for i := mantissaStart + firstDigit; i < mantissaEnd; i += 3 {
			buf.WriteRune(',')
			buf.WriteString(s[i : i+3])
		}
	}
	buf.WriteString(s[mantissaEnd:])
	return buf.String()
}

// !-
