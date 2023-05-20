package testSomething

import (
	"fmt"
	inputNumberText "spaghetti-code/number_in_text_un_scramble/numbertext"
	"testing"

	exSlices "golang.org/x/exp/slices"
)

func TestNumberInTextUnscramble(t *testing.T) {
	var d inputNumberText.NumberText
	have := "nnsoitnexniteeeonvswe"
	d.SetInput(have)
	d.NumberIntextUnscramble()

	numberListResult := TestingSuiteNumberList(d.NumberList)
	if numberListResult != "" {
		t.Error(numberListResult)
	}
	textNumberListResult := TestingSuiteTextNumberList(d.TextNumberList)
	if textNumberListResult != "" {
		t.Error(textNumberListResult)
	}
}

func TestingSuiteNumberList(testInputNumberList []int) string {
	got := testInputNumberList
	wantList := [...]int{1, 9, 10, 2, 6, 7}
	if !(len(got) == len(wantList)) {
		return "element length are not equal"
	}

	for _, want := range wantList {
		if exSlices.Index(got, want) == -1 {
			return fmt.Sprint("some element is in correct.", want, "is expected")
		}
	}
	return ""
}

func TestingSuiteTextNumberList(testInputTextNumberList []string) string {
	got := testInputTextNumberList
	wantList := [...]string{"one", "nine", "ten", "two", "six", "seven"}
	if !(len(got) == len(wantList)) {
		return fmt.Sprintln("element length are not equal.", "expected:", len(wantList), "got:", len(got))
	}

	for _, want := range wantList {
		if exSlices.Index(got, want) == -1 {
			return fmt.Sprint("some element is in correct.", want, "is expected")
		}
	}

	return ""
}
