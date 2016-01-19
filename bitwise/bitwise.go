package bitwise

import (
	"math"
)

//XOR takes two equal length bytes arrays and XORS them

func XOR(b []byte, c []byte) []byte {
	n := len(b)
	o := make([]byte, n)
	for i := 0; i < n; i++ {
		o[i] = b[i] ^ c[i]
	}
	return (o)
}

//Single Byte takes a byte , and a byte array. Loops over the array XORing with byte

func SBXOR(b []byte, c byte) []byte {
	n := len(b)
	o := make([]byte, n)
	for i := 0; i < n; i++ {
		o[i] = c ^ b[i]
	}
	return (o)
}

func BRUTEXOR(b []byte) [][]byte {

	o := [][]byte{}

	for i := '\x00'; i <= '\xFF'; i++ {
		o = append(o, SBXOR(b, byte(i)))
	}

	return (o)
}

//repeating key XOR

func RXOR(b []byte, c []byte) []byte {
	n := len(b)
	o := make([]byte, n)

	j := 0

	for i := 0; i < n; i++ {

		o[i] = b[i] ^ c[j]

		j++

		if j == (len(c) - 1) {
			j = 0
		}

	}

	return (o)
}

func HAM(b []byte, c []byte) int {

	n := len(b)
	var j int
	var o byte

	j = 0

	for i := 0; i < n; i++ {

		o = b[i] ^ c[i]

		j += COUNTBITS(o)

	}

	return (j)

}

func COUNTBITS(b byte) int {

	var j int

	j = 0

	if b >= 128 {
		b = byte(math.Mod(float64(b), 128))
		j++
	}

	if b >= 64 {
		b = byte(math.Mod(float64(b), 64))
		j++
	}

	if b >= 32 {
		b = byte(math.Mod(float64(b), 32))
		j++
	}

	if b >= 16 {
		b = byte(math.Mod(float64(b), 16))
		j++
	}

	if b >= 8 {
		b = byte(math.Mod(float64(b), 8))
		j++
	}

	if b >= 4 {
		b = byte(math.Mod(float64(b), 4))
		j++
	}

	if b >= 2 {
		b = byte(math.Mod(float64(b), 2))
		j++
	}

	if b >= 1 {
		b = byte(math.Mod(float64(b), 1))
		j++
	}
	return j
}
