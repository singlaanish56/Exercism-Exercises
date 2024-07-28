package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	
	return fmt.Sprintf("This is the number %.1f",f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	
	return fmt.Sprintf("This is a box containing the number %.1f", float64(nb.Number()))

}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	
	var x string
	
	switch fnb.(type){
	case FancyNumber:
		x = fnb.Value()
	default:
		x = "0"
	}

	i, _ := strconv.Atoi(x)

	return i
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	
	return fmt.Sprintf("This is a fancy box containing the number %0.1f",float64(ExtractFancyNumber(fnb)))
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	
	var ans string

	switch v :=i.(type){
	case int:
		ans = DescribeNumber(float64(v))
	case float64:
		ans = DescribeNumber(v)
	case NumberBox:
		ans = DescribeNumberBox(v)
	case FancyNumberBox:
		ans = DescribeFancyNumberBox(v)
	default:
		ans ="Return to sender"
	}

	return ans
}
