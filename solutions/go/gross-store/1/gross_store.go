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
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	value, hasUnit := units[unit]
	if !hasUnit {
		return hasUnit
	}
	bill[item] += value
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	if _, hasItem := bill[item]; !hasItem {
		return hasItem
	}
	if _, hasUnit := units[unit]; !hasUnit {
		return hasUnit
	}
	if bill[item]-units[unit] < 0 {
		return false
	}
	if bill[item]-units[unit] == 0 {
		delete(bill, item)
		return true
	}
	bill[item] -= units[unit]
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	if _, hasItem := bill[item]; !hasItem {
		return 0, hasItem
	}
	return bill[item], true
}
