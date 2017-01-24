package letter

import "sync"

// FreqMap maps letters to their frequencies
type FreqMap map[rune]int

// Frequency counts the frequency of the letters
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency fans out processing of strings and fans in the counts
func ConcurrentFrequency(ss []string) FreqMap {
	wg := sync.WaitGroup{}
	in := make(chan FreqMap)
	out := make(chan FreqMap)

	// Fan out
	wg.Add(len(ss))
	for ix := range ss {
		go func(s string) {
			in <- Frequency(s)
			wg.Done()
		}(ss[ix])
	}

	go func() {
		wg.Wait()
		close(in)
	}()

	// Fan in
	go func() {
		totals := make(FreqMap)
		for fm := range in {
			for k, v := range fm {
				totals[k] += v
			}
		}
		out <- totals
	}()

	return <-out
}
