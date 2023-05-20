package main

import (
	"fmt"
	NumberTextPackage "spaghetti-code/number_in_text_un_scramble/numbertext"
)

// scrambled from oneninetentwosixseven
// expecting array similar to
// [one, nine, ten, two, six, seven]
// OR
// [1, 9, 10, 2, 6, 7]
func main() {
	fmt.Println("number in text unscramble")
	d := NumberTextPackage.NumberText{}
	const input = "nnsoitnexniteeeonvswe"
	d.SetInput(input)
	d.NumberIntextUnscramble()
	fmt.Println("string input: ", d.Input)
	fmt.Println("text number list: ", d.TextNumberList)
	fmt.Println("number list: ", d.NumberList)
}
