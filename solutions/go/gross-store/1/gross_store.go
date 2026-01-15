package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
        "quarter_of_a_dozen" : 3,
        "half_of_a_dozen" : 6,
        "dozen" : 12,
        "small_gross" : 120,
        "gross" : 144,
        "great_gross" : 1728,
    }
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	_, ok := units[unit]
    if !ok {
        return false
    }
    _, okk := bill[item]
    if !okk {
        bill[item] = units[unit]
        return true
    } else {
        bill[item] += units[unit]
        return true
    }
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	u, ok := units[unit]
    if !ok {
        return false
    }
    b, okk := bill[item]
    if !okk {
        return false
    } else if b < u {
        return false
    } else if b == u {
        delete(bill, item)
        return true
    } else {
        bill[item] -= u
        return true
    }
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	b, ok := bill[item]
    if !ok {
        return 0, false
    } else {
        return b, true
    }
}
