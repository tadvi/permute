/* Package for permutations and combinations */
package permute

// Ints returns 0..n slice as permutations of slices of int.
// For example n=2 becomes [0, 1] [1, 0].
func Ints(n int, fn func([]int)) {
	a := make([]int, n)
	for i := range a {
		a[i] = i
	}
	permutations(a, 0, fn)
}

// Bools returns 0..n slice as permutations of slices of bool.
// For example n=2 becomes [false, false] [true, false], [false, true] [true, true].
func Bools(n int, fn func([]bool)) {
	if n > 64 {
		panic("permute: Bools permutations limited to 64 values")
	}
	var max uint64 = 1 << uint(n)
	var i uint64
	a := make([]bool, n)

	for ; i < max; i++ {
		var j uint
		for ; j < uint(n); j++ {
			if (i & (1 << j)) > 0 {
				a[j] = true
			} else {
				a[j] = false
			}
		}
		fn(a)
	}
}

// Combinations generates all combinations of a slice of int values.
func Combinations(a []int, fn func([]int)) {
	combinations(a, 0, fn)
}

// SliceInts generates all permutations of a slice of int values.
func SliceInts(a []int, fn func([]int)) {
	permutations(a, 0, fn)
}

// SliceBytes generates all permutations of a slice of byte values.
func SliceBytes(a []byte, fn func([]byte)) {
	permutationsBytes(a, 0, fn)
}

// Tuples returns all slice values in pairs.
// For example [0, 1, 2] becomes [0, 1] [0, 2] [1, 2].
func Tuples(a []int, fn func([]int)) {
	p := make([]int, 2)
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			p[0], p[1] = a[i], a[j]
			fn(p)
		}
	}
}

// Triples returns all slice values in groups of triples.
// For example [0, 1, 2, 3] becomes [0, 1, 2], [0, 1, 3], [0, 2, 3], [1, 2, 3].
func Triples(a []int, fn func([]int)) {
	p := make([]int, 3)
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			for k := j + 1; k < len(a); k++ {
				p[0], p[1], p[2] = a[i], a[j], a[k]
				fn(p)
			}
		}
	}
}

func permutations(a []int, start int, fn func([]int)) {
	if start == len(a)-1 {
		// bottom of recursion. You get here once for each permutation.
		fn(a)
		return
	}
	for i := start; i < len(a); i++ {
		a[i], a[start] = a[start], a[i]
		permutations(a, start+1, fn)
		a[i], a[start] = a[start], a[i]
	}
}

func permutationsBytes(a []byte, start int, fn func([]byte)) {
	if start == len(a)-1 {
		// bottom of recursion. You get here once for each permutation.
		fn(a)
		return
	}
	for i := start; i < len(a); i++ {
		a[i], a[start] = a[start], a[i]
		permutationsBytes(a, start+1, fn)
		a[i], a[start] = a[start], a[i]
	}
}

func combinations(a []int, start int, fn func([]int)) {
	if start > 0 {
		fn(a[:start])
	}

	if start == len(a)-1 {
		// bottom of recursion. You get here once for each permutation.
		return
	}

	for i := start; i < len(a); i++ {
		a[i], a[start] = a[start], a[i]
		combinations(a, start+1, fn)
		a[i], a[start] = a[start], a[i]
	}
}
