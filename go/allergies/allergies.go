// Package allergies tests for common allergies
package allergies

var testVersion = 1

// Map of allergy name to bit offset
var testedAllergies = map[string]uint{
	"eggs":         0,
	"peanuts":      1,
	"shellfish":    2,
	"strawberries": 3,
	"tomatoes":     4,
	"chocolate":    5,
	"pollen":       6,
	"cats":         7,
}

// isAllergic extracts the common test logic for performance reasons
//
// Simply calling AllergicTo() from Allergies() adds 50% to the execution time
// due to the extra test for existence of the key, even though we know the
// key is there and what the bit number is.
func isAllergic(score, bit uint) bool {
	return score>>bit&1 == 1
}

// Allergies returns a list of allergies based on a score
func Allergies(score uint) (result []string) {
	for allergen, bit := range testedAllergies {
		if isAllergic(score, bit) {
			result = append(result, allergen)
		}
	}
	return
}

// AllergicTo returns true if the specified allergy was found
func AllergicTo(score uint, allergen string) (result bool) {
	if bit, ok := testedAllergies[allergen]; ok {
		result = isAllergic(score, bit)
	}
	return
}
