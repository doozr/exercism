package pythagorean

import "sort"

// Triplet of pythagorean triplet values
type Triplet [3]int

func newTriplet(a, b, c int) Triplet {
	s := []int{a, b, c}
	sort.Ints(s)
	return Triplet{s[0], s[1], s[2]}
}

func (t Triplet) sum() int {
	return t[0] + t[1] + t[2]
}

func (t Triplet) multiply(x int) Triplet {
	return newTriplet(t[0]*x, t[1]*x, t[2]*x)
}

func (t Triplet) withinBounds(min, max int) bool {
	return t[0] >= min && t[0] <= max &&
		t[1] >= min && t[1] <= max &&
		t[2] >= min && t[2] <= max
}

// TripletSorter allows lexicographic sorting of Triples
type TripletSorter []Triplet

func (ts TripletSorter) Len() int {
	return len(ts)
}

func (ts TripletSorter) Swap(i, j int) {
	ts[i], ts[j] = ts[j], ts[i]
}

func (ts TripletSorter) Less(i, j int) bool {
	return ts[i][0] < ts[j][0]
}

func mn(limit int) (triplets []Triplet) {
	triplets = make([]Triplet, 0)
	for m := 1; m < limit; m++ {
		for n := 1; n < m; n++ {
			a := (m * m) - (n * n)
			b := 2 * m * n
			c := (m * m) + (n * n)
			triplets = append(triplets, newTriplet(a, b, c))
		}
	}
	return
}

// Sum returns all triplets with a given perimeter
func Sum(num int) (triplets []Triplet) {
	tripletSet := make(map[Triplet]bool)
	for _, triplet := range mn(num / 2) {
		total := triplet.sum()

		if total == num {
			tripletSet[triplet] = true
		}

		if total < num {
			for x := 2; total <= num; x++ {
				scaledTriple := triplet.multiply(x)
				total = scaledTriple.sum()
				if total == num {
					tripletSet[scaledTriple] = true
				}
			}
		}
	}

	triplets = make([]Triplet, 0)
	for k := range tripletSet {
		triplets = append(triplets, k)
	}

	sort.Sort(TripletSorter(triplets))
	return
}

// Range returns all triplets with sides in the given range
func Range(min, max int) (triplets []Triplet) {
	tripletSet := make(map[Triplet]bool)
	for _, triplet := range mn(max / 2) {
		if triplet.withinBounds(min, max) {
			tripletSet[triplet] = true
		}
	}

	triplets = make([]Triplet, 0)
	for k := range tripletSet {
		triplets = append(triplets, k)
	}

	sort.Sort(TripletSorter(triplets))
	return
}
