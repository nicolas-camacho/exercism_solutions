// Package leap works within a year to solve the leap problem.
package leap

// IsLeapYear recive a year and determine if it's a leap year or not.
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
