package main

import (
	"fmt"
	"github.com/scripterio/go-cryptopals/bitwise"
	"github.com/scripterio/go-cryptopals/hexstring"
)

func main() {

	input := "482169607764216021657364606c"

	raw := hexstring.Decode(input)

	bruteout := bitwise.BRUTEXOR(raw)

	_, key := hexstring.ScoreEnglish(bruteout)

	output := string(bruteout[key])

	//Test
	fmt.Printf("Set 1   Challenge 3 \n")

	if output == "I have a dream" {

		fmt.Printf("PASS!!\n\n")

	} else {
		fmt.Printf("##FAIL##\n\n")
	}

}
