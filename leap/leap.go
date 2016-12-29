/*
Package leap provides leap year calculation function
*/
package leap

// testVersion should match the targetTestVersion in the test file.
const testVersion = 2

// Returns true when year is a leap year
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
