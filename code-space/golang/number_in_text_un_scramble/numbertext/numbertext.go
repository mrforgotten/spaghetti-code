package numbertext

import "strings"

type NumberText struct {
	Input          string
	TextNumberList []string
	NumberList     []int
}

func (d *NumberText) SetInput(input string) {
	d.Input = input
	d.TextNumberList = []string{}
	d.NumberList = []int{}
}

func (d *NumberText) NumberIntextUnscramble() {
	var dInput string = d.Input
	for len(dInput) >= 3 {

		if strings.Contains(dInput, "w") {
			dInput = strings.Replace(
				strings.Replace(
					strings.Replace(dInput, "t", "", 1), "w", "", 1), "o", "", 1)
			d.NumberList = append(d.NumberList, 2)
			d.TextNumberList = append(d.TextNumberList, "two")

		} else if strings.Contains(dInput, "u") {
			dInput =
				strings.Replace(
					strings.Replace(
						strings.Replace(
							strings.Replace(dInput, "f", "", 1), "o", "", 1), "u", "", 1), "r", "", 1)
			d.NumberList = append(d.NumberList, 4)
			d.TextNumberList = append(d.TextNumberList, "four")

		} else if strings.Contains(dInput, "x") {
			dInput = strings.Replace(
				strings.Replace(
					strings.Replace(dInput, "s", "", 1), "i", "", 1), "x", "", 1)
			d.NumberList = append(d.NumberList, 6)
			d.TextNumberList = append(d.TextNumberList, "six")

		} else if strings.Contains(dInput, "g") {
			dInput =
				strings.Replace(
					strings.Replace(
						strings.Replace(
							strings.Replace(
								strings.Replace(
									dInput, "t", "", 1), "h", "", 1), "g", "", 1), "i", "", 1), "e", "", 1)
			d.NumberList = append(d.NumberList, 8)
			d.TextNumberList = append(d.TextNumberList, "eight")

		} else if strings.Contains(dInput, "z") {
			dInput =
				strings.Replace(
					strings.Replace(
						strings.Replace(
							strings.Replace(dInput, "o", "", 1), "r", "", 1), "e", "", 1), "z", "", 1)
			d.NumberList = append(d.NumberList, 0)
			d.TextNumberList = append(d.TextNumberList, "zero")

		} else if strings.Contains(dInput, "o") {
			dInput =
				strings.Replace(strings.Replace(strings.Replace(dInput, "e", "", 1), "n", "", 1), "o", "", 1)
			d.NumberList = append(d.NumberList, 1)
			d.TextNumberList = append(d.TextNumberList, "one")

		} else if strings.Contains(dInput, "h") {
			dInput =
				strings.Replace(
					strings.Replace(
						strings.Replace(
							strings.Replace(
								dInput, "e", "", 2), "r", "", 1), "h", "", 1), "t", "", 1)
			d.NumberList = append(d.NumberList, 3)
			d.TextNumberList = append(d.TextNumberList, "three")

		} else if strings.Contains(dInput, "f") {
			dInput =
				strings.Replace(
					strings.Replace(
						strings.Replace(
							strings.Replace(
								dInput, "e", "", 1), "v", "", 1), "i", "", 1), "f", "", 1)
			d.NumberList = append(d.NumberList, 5)
			d.TextNumberList = append(d.TextNumberList, "five")

		} else if strings.Contains(dInput, "s") {
			dInput =
				strings.Replace(
					strings.Replace(
						strings.Replace(
							strings.Replace(
								dInput, "n", "", 1), "v", "", 1), "e", "", 2), "s", "", 1)
			d.NumberList = append(d.NumberList, 7)
			d.TextNumberList = append(d.TextNumberList, "seven")

		} else if strings.Contains(dInput, "n") && strings.Contains(dInput, "i") && strings.Contains(dInput, "e") {
			dInput =
				strings.Replace(
					strings.Replace(
						strings.Replace(dInput, "e", "", 1), "i", "", 1), "n", "", 2)
			d.NumberList = append(d.NumberList, 9)
			d.TextNumberList = append(d.TextNumberList, "nine")

		} else if strings.Contains(dInput, "t") && strings.Contains(dInput, "e") && strings.Contains(dInput, "n") {
			dInput =
				strings.Replace(
					strings.Replace(
						strings.Replace(dInput, "t", "", 1), "e", "", 1), "n", "", 1)
			d.NumberList = append(d.NumberList, 10)
			d.TextNumberList = append(d.TextNumberList, "ten")

		}
	}
}
