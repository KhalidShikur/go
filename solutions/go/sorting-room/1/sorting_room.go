package sorting
import (
    "fmt"
    "strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	f := float64(nb.Number())
    return fmt.Sprintf("This is a box containing the number %.1f", f)
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
    result := 0
    switch fnb.(type) {
        case FancyNumber:
            i, _ := strconv.Atoi(fnb.Value())
            result = i
        default:
            result = 0
    }
    
    return result
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
    extract := 0
	switch fnb.(type) {
        case FancyNumber:
            extract = ExtractFancyNumber(fnb)
            break
        default: 
            extract = 0
    }

    return fmt.Sprintf("This is a fancy box containing the number %.1f", float64(extract))
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i any) string {
    result := "Return to sender"
	switch i.(type) {
        case int:
            f, _ := i.(int)
            result = DescribeNumber(float64(f))
        case float64:
            f, _ := i.(float64)
            result = DescribeNumber(f)
        case NumberBox:
            f, _ := i.(NumberBox)
            result = DescribeNumberBox(f)
        case FancyNumberBox:
            f, _ := i.(FancyNumberBox)
            result = DescribeFancyNumberBox(f)
        default:
            result = "Return to sender"
    }

    return result
}
