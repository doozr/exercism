// Leap calculates leap years
package leap

// testVersion should match the targetTestVersion in the test file.
const testVersion = 2

// IsLeapYear calculates whether a year is a leap year or not.
func IsLeapYear(year int) bool {
	return year%400 == 0 || (year%4 == 0 && year%100 != 0)
}
