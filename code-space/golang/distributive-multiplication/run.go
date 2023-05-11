package main

import (
	"fmt"
	"strings"
)

// scrambled from oneninetentwosixseven
// expecting array similar to
// [one, nine, ten, two, six, seven]
// OR
// [1, 9, 10, 2, 6, 7]
func main() {
	fmt.Println("distributive multiplication")
	var d Distributive
	const input = "nnsoitnexniteeeonvswe"
	d.SetInput(input)
	d.DistributiveMultiplication()
	fmt.Println("string input: ", d.input)
	fmt.Println("text number list: ", d.text_number_list)
	fmt.Println("number list: ", d.number_list)
}

type Distributive struct {
	input            string
	text_number_list []string
	number_list      []int
}

func (d *Distributive) SetInput(input string) {
	d.input = input
	d.text_number_list = []string{}
	d.number_list = []int{}
}

func (d *Distributive) DistributiveMultiplication() {
	var dInput string = d.input
	for len(dInput) >= 3 {

		if strings.Contains(dInput, "w") {
			dInput = strings.Replace(
				strings.Replace(
					strings.Replace(dInput, "t", "", 1), "w", "", 1), "o", "", 1)
			d.number_list = append(d.number_list, 2)
			d.text_number_list = append(d.text_number_list, "two")

		} else if strings.Contains(dInput, "u") {
			dInput =
				strings.Replace(
					strings.Replace(
						strings.Replace(
							strings.Replace(dInput, "f", "", 1), "o", "", 1), "u", "", 1), "r", "", 1)
			d.number_list = append(d.number_list, 4)
			d.text_number_list = append(d.text_number_list, "four")

		} else if strings.Contains(dInput, "x") {
			dInput = strings.Replace(
				strings.Replace(
					strings.Replace(dInput, "s", "", 1), "i", "", 1), "x", "", 1)
			d.number_list = append(d.number_list, 6)
			d.text_number_list = append(d.text_number_list, "six")

		} else if strings.Contains(dInput, "g") {
			dInput =
				strings.Replace(
					strings.Replace(
						strings.Replace(
							strings.Replace(
								strings.Replace(
									dInput, "t", "", 1), "h", "", 1), "g", "", 1), "i", "", 1), "e", "", 1)
			d.number_list = append(d.number_list, 8)
			d.text_number_list = append(d.text_number_list, "eight")

		} else if strings.Contains(dInput, "z") {
			dInput =
				strings.Replace(
					strings.Replace(
						strings.Replace(
							strings.Replace(dInput, "o", "", 1), "r", "", 1), "e", "", 1), "z", "", 1)
			d.number_list = append(d.number_list, 0)
			d.text_number_list = append(d.text_number_list, "zero")

		} else if strings.Contains(dInput, "o") {
			dInput =
				strings.Replace(strings.Replace(strings.Replace(dInput, "e", "", 1), "n", "", 1), "o", "", 1)
			d.number_list = append(d.number_list, 1)
			d.text_number_list = append(d.text_number_list, "one")

		} else if strings.Contains(dInput, "h") {
			dInput =
				strings.Replace(
					strings.Replace(
						strings.Replace(
							strings.Replace(
								dInput, "e", "", 2), "r", "", 1), "h", "", 1), "t", "", 1)
			d.number_list = append(d.number_list, 3)
			d.text_number_list = append(d.text_number_list, "three")

		} else if strings.Contains(dInput, "f") {
			dInput =
				strings.Replace(
					strings.Replace(
						strings.Replace(
							strings.Replace(
								dInput, "e", "", 1), "v", "", 1), "i", "", 1), "f", "", 1)
			d.number_list = append(d.number_list, 5)
			d.text_number_list = append(d.text_number_list, "five")

		} else if strings.Contains(dInput, "s") {
			dInput =
				strings.Replace(
					strings.Replace(
						strings.Replace(
							strings.Replace(
								dInput, "n", "", 1), "v", "", 1), "e", "", 2), "s", "", 1)
			d.number_list = append(d.number_list, 7)
			d.text_number_list = append(d.text_number_list, "seven")

		} else if strings.Contains(dInput, "n") && strings.Contains(dInput, "i") && strings.Contains(dInput, "e") {
			dInput =
				strings.Replace(
					strings.Replace(
						strings.Replace(dInput, "e", "", 1), "i", "", 1), "n", "", 2)
			d.number_list = append(d.number_list, 9)
			d.text_number_list = append(d.text_number_list, "nine")

		} else if strings.Contains(dInput, "t") && strings.Contains(dInput, "e") && strings.Contains(dInput, "n") {
			dInput =
				strings.Replace(
					strings.Replace(
						strings.Replace(dInput, "t", "", 1), "e", "", 1), "n", "", 1)
			d.number_list = append(d.number_list, 10)
			d.text_number_list = append(d.text_number_list, "ten")

		}
	}
}
