package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"

	"gopl.io/ch2/lenconv"
	"gopl.io/ch2/massconv"
	"gopl.io/ch2/tempconv"
)

const (
	temp   = "t"
	length = "l"
	mass   = "m"
)

func main() {
	var typeFlag = flag.String("type", "", "type of measurement\npossible values:\n - t — temperature\n - l — length\n - m — mass")
	flag.Parse()
	if !typeFlagValid(typeFlag) {
		fmt.Fprintf(os.Stderr, "Unexpected type value: %s\n", *typeFlag)
		os.Exit(1)
	}
	if len(flag.Args()) == 0 {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			if input.Text() == "quit" {
				break
			}
			conv(input.Text(), typeFlag)
		}
	}
	for _, arg := range flag.Args() {
		conv(arg, typeFlag)
	}
}

func conv(arg string, typeFlag *string) {
	val, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		os.Exit(1)
	}
	switch *typeFlag {
	case temp:
		c := tempconv.Celsius(val)
		f := tempconv.Fahrenheit(val)
		k := tempconv.Kelvin(val)
		fmt.Printf("%s = %s, %s = %s\n%s = %s, %s = %s\n%s = %s, %s = %s\n\n",
			c, tempconv.CToF(c), c, tempconv.CToK(c),
			f, tempconv.FToC(f), f, tempconv.FToK(f),
			k, tempconv.KToC(k), k, tempconv.KToF(k))
	case length:
		m := lenconv.Meter(val)
		f := lenconv.Foot(val)
		fmt.Printf("%s = %s, %s = %s\n\n",
			m, lenconv.MTof(m), f, lenconv.FToM(f))
	case mass:
		kg := massconv.Kilogram(val)
		p := massconv.Pound(val)
		fmt.Printf("%s = %s, %s = %s\n\n",
			kg, massconv.KToP(kg), p, massconv.PToK(p))
	}
}

func typeFlagValid(typeFlag *string) bool {
	return *typeFlag == temp ||
		*typeFlag == length ||
		*typeFlag == mass
}
