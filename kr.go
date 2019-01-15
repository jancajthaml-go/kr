//
// Information about the algorithm is available on Wikipedia
//
// https://en.wikipedia.org/wiki/Rabin%E2%80%93Karp_algorithm
//
package kr

func Search(input, pattern string) []int {
	var M = len(pattern)
	var N = len(input)

	if N == 0 || M == 0 {
		return nil
	}

	var result = make([]int, 0)
	var i int
	var j int
	var p = 0
	var t = 0
	var h = 1
	var prime = 2017

	for i = 0; i < M-1; i++ {
		h = (h << 9) % prime
		p = (p<<9 + int(pattern[i])) % prime
		t = (t<<9 + int(input[i])) % prime
	}

	p = (p<<9 + int(pattern[i])) % prime
	t = (t<<9 + int(input[i])) % prime

	for i = 0; i <= N-M; i++ {
		if p == t {
			for j = 0; j < M; j++ {
				if input[i+j] != pattern[j] {
					break
				}
			}

			if j == M {
				result = append(result, i)
			}
		}

		if i < N-M {
			t = ((t-int(input[i])*h)<<9 + int(input[i+M])) % prime

			if t < 0 {
				t = t + prime
			}
		}
	}

	return result
}