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

// Allergies returns a list of allergies based on a score
func Allergies(score uint) (result []string) {
	for allergen, bit := range testedAllergies {
		if score>>bit&1 == 1 {
			result = append(result, allergen)
		}
	}
	return
}

// AllergicTo returns true if the specified allergy was found
func AllergicTo(score uint, allergen string) bool {
	if bit, ok := testedAllergies[allergen]; ok {
		return score>>bit&1 == 1
	}
	return false
}
