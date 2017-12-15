// http://adventofcode.com/2017/day/15

package main

import "fmt"

func main() {
	// Test values (Answers: part1 - 588, part2 - 309)
	// a := 65
	// b := 8921

	a := 679
	b := 771
	factorA := 16807
	factorB := 48271
	divider := 2147483647
	cycles1 := 40000000
	cycles2 := 5000000

	fmt.Println(generator1(a, b, factorA, factorB, divider, cycles1))
	fmt.Println(generator2(a, b, factorA, factorB, divider, cycles2))
}

func generator1(a, b, fa, fb, div, c int) int {
	count := 0
	for i := 0; i < c; i++ {
		a = (a * fa) % div
		b = (b * fb) % div
		if checkMatch(a, b) {
			count++
		}
	}
	return count
}

func generator2(a, b, fa, fb, div, c int) int {
	resultsA := []int{}
	resultsB := []int{}

	count := 0
	pairCount := 0
	for pairCount < c {
		a = (a * fa) % div
		b = (b * fb) % div

		if a%4 == 0 {
			resultsA = append(resultsA, a)
		}
		if b%8 == 0 {
			resultsB = append(resultsB, b)
		}
		if len(resultsA) > 0 && len(resultsB) > 0 {
			pairCount++
			a1, aRem := resultsA[0], resultsA[1:]
			resultsA = aRem
			b1, bRem := resultsB[0], resultsB[1:]
			resultsB = bRem

			if checkMatch(a1, b1) {
				count++
			}
		}
	}
	return count
}

func checkMatch(n1 int, n2 int) bool {
	h1 := fmt.Sprintf("%x", n1)
	h1 = padZeros(h1, 4)
	hs1 := h1[len(h1)-4:]

	h2 := fmt.Sprintf("%x", n2)
	h2 = padZeros(h2, 4)
	hs2 := h2[len(h2)-4:]

	return hs1 == hs2
}

func padZeros(s string, n int) string {
	for len(s) < n {
		s = "0" + s
	}
	return s
}
