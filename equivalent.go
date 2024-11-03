package gloop

import "iter"

// Equivalent checks if two given [iter.Seq] sequences are equal in
// contents, ignoring order.
func Equivalent[V comparable](seq1 iter.Seq[V], seq2 iter.Seq[V]) bool {
	next1, stop1 := iter.Pull(seq1)
	defer stop1()

	next2, stop2 := iter.Pull(seq2)
	defer stop2()

	m1 := make(map[V]int)
	m2 := make(map[V]int)

	for {
		value1, ok1 := next1()
		value2, ok2 := next2()

		if !ok1 && !ok2 {
			break
		}

		if ok1 != ok2 {
			return false
		}

		if value1 == value2 {
			continue
		}

		n12, ok12 := m1[value2]
		if ok12 {
			if n12 > 1 {
				m1[value2] = n12 - 1
			} else {
				delete(m1, value2)
			}
		}

		n21, ok21 := m2[value1]
		if ok21 {
			if n21 > 1 {
				m2[value1] = n21 - 1
			} else {
				delete(m2, value1)
			}
		}

		if !ok21 {
			n11, ok11 := m1[value1]
			if ok11 {
				m1[value1] = n11 + 1
			} else {
				m1[value1] = 1
			}
		}

		if !ok12 {
			n22, ok22 := m2[value2]
			if ok22 {
				m2[value2] = n22 + 1
			} else {
				m2[value2] = 1
			}
		}
	}

	return len(m1) == 0 && len(m2) == 0
}

// Equivalent2 checks if two given [iter.Seq2] sequences are equal in
// contents, ignoring order.
func Equivalent2[K, V comparable](seq1 iter.Seq2[K, V], seq2 iter.Seq2[K, V]) bool {
	return Equivalent(KeyValue2(seq1), KeyValue2(seq2))
}
