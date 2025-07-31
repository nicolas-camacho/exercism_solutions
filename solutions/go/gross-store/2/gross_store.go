package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	value, hasUnit := units[unit]
	if !hasUnit {
		return false
	}
	bill[item] += value
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	if _, hasItem := bill[item]; !hasItem {
		return false
	}
	if _, hasUnit := units[unit]; !hasUnit {
		return false
	}

	remaining := bill[item] - units[unit]

	switch {
	case remaining < 0:
		return false
	case remaining == 0:
		delete(bill, item)
		return true
	default:
		bill[item] -= units[unit]
	}

	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	if _, hasItem := bill[item]; !hasItem {
		return 0, false
	}
	return bill[item], true
}
