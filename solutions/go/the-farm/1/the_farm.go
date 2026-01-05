package thefarm
import (
    "fmt"
    "errors"
)

// TODO: define the 'DivideFood' function
func DivideFood(fcalc FodderCalculator, n int) (float64, error) {
    fodderAmount, err := fcalc.FodderAmount(n)
    if err != nil {
        return 0, err
    }
    fatteningFactor, err := fcalc.FatteningFactor()
    if err != nil {
        return 0, err
    }

    return ((fodderAmount/float64(n)) * fatteningFactor), nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fcalc FodderCalculator, n int) (float64, error) {
    if n <= 0 {
        return 0, errors.New("invalid number of cows")
    } else {
        return DivideFood(fcalc, n)
    }
}

// TODO: define the 'ValidateNumberOfCows' function
type InvalidCowsError struct {
    numberOfCows int
    message string
}

func (e InvalidCowsError) Error() string {
    return fmt.Sprintf("%d cows are invalid: %s", e.numberOfCows, e.message)
}

func ValidateNumberOfCows(n int) error {
    if n < 0 {
        return &InvalidCowsError{
            numberOfCows : n,
            message : "there are no negative cows",
        }
    } else if n == 0 {
        return &InvalidCowsError{
            numberOfCows : n,
            message : "no cows don't need food",
        }
    } else {
        return nil
    }
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
