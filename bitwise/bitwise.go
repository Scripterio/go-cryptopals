package bitwise

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
