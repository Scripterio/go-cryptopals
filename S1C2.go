package main

import (
	"fmt"
	"github.com/scripterio/go-cryptopals/bitwise"
	"github.com/scripterio/go-cryptopals/hexstring"
)

func main() {

	input1 := "f0f0f0"
	input2 := "0f0f0f"

	raw1 := hexstring.Decode(input1)
	raw2 := hexstring.Decode(input2)

	xord := bitwise.XOR(raw1, raw2)

	output := hexstring.Encode(xord)

	fmt.Printf("Set 1   Challenge 2 \n")

	//Test
	if output == "ffffff" {

		fmt.Printf("PASS!!\n\n")

	} else {
		fmt.Printf("##FAIL##\n\n")
	}

}
