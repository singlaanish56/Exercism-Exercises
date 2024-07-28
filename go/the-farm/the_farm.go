package thefarm

import (
	"errors"
	"fmt"

)

// TODO: define the 'DivideFood' function
func DivideFood(fc FodderCalculator, nc int) (float64, error) {
	totalAmt, err := fc.FodderAmount(nc)
	if err != nil{
		return 0, err
	}
	
	f, e := fc.FatteningFactor()
	if e != nil{
		return 0, e
	}

	return float64(totalAmt * f) / float64(nc), nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fc FodderCalculator, nc int) (float64, error) {

	if(nc>0){
		return DivideFood(fc, nc)
	}

	return 0, errors.New("invalid number of cows")
}

type InvalidCowsError struct{
	nc int
	message string
}

func (i *InvalidCowsError) Error() string{
	return fmt.Sprintf("%d cows are invalid: %s",i.nc, i.message)
}
// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(nc int) error{
	if(nc <0){
		return &InvalidCowsError{
			nc,
			"there are no negative cows",
		}
	}else if(nc ==0){
		return &InvalidCowsError{
			nc,
			"no cows don't need food",
		}
	}

	return nil
}
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
