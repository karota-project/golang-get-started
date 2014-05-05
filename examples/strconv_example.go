package main

import (
	"fmt"
	"strconv"
)

func byteToString(c []byte) string {
	n := -1
	for i, b := range c {
		if b == 0 {
			break
		}
		n = i
	}
	return string(c[:n+1])
}

func main() {
	/* Append */
	s := make([]byte, 0, 100)
	/* func AppendInt(dst []byte, i int64, base int) []byte */
	s = strconv.AppendInt(s, 123, 10)
	/* func AppendBool(dst []byte, b bool) []byte */
	s = strconv.AppendBool(s, false)

	fmt.Println("byte : ", s, " conversion to string : ", byteToString(s[:]))

	/* Format */
	/* func FormatBool(b bool) string */
	b := strconv.FormatBool(false)
	/* func Itoa(i int) string */
	i := strconv.Itoa(456)

	fmt.Println(b, i)

	/* Parse */
	/* func Atoi(s string) (i int, err error) */
	a, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println(err)
	}
	/* func ParseFloat(s string, bitSize int) (f float64, err error) */
	f, err := strconv.ParseFloat("123.23", 64)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(a, f)
}

/*
$ go run strconv_example.go
byte :  [49 50 51 102 97 108 115 101]
conversion to string :  123false
false 456
123 123.23
*/
