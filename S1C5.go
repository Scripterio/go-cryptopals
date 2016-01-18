package main

import (
	"fmt"

	"github.com/scripterio/go-cryptopals/bitwise"
	"github.com/scripterio/go-cryptopals/hexstring"
)

func main() {

	input := "We hold these truths to be self evident: that all men are created equal."

	cipherkey := "MLK"

	rawinput := []byte(input)

	rawkey := []byte(cipherkey)

	ciphertext := bitwise.RXOR(rawinput, rawkey)

	output := hexstring.Encode(ciphertext)

	//Test
	fmt.Printf("Set 1   Challenge 5 \n")

	if output == "1a296d242220296c3924283f286c393e3838253f6d38226c2f296d3f28202b6c283a2428282239766d38252d396c2c20216c2029236c2c3e286c2e3e282d3929296c283d382d2162" {

		fmt.Printf("PASS!!\n\n")

	} else {
		fmt.Printf("##FAIL##\n\n")
	}

}
