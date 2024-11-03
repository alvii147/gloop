package gloop

import "iter"

// Equal checks if two given [iter.Seq] sequences are exactly equal in
// contents and order.
func Equal[V comparable](seq1 iter.Seq[V], seq2 iter.Seq[V]) bool {
	next1, stop1 := iter.Pull(seq1)
	defer stop1()

	next2, stop2 := iter.Pull(seq2)
	defer stop2()

	for {
		value1, ok1 := next1()
		value2, ok2 := next2()

		if !ok1 && !ok2 {
			break
		}

		if ok1 != ok2 {
			return false
		}

		if value1 != value2 {
			return false
		}
	}

	return true
}

// Equal2 checks if two given [iter.Seq2] sequences are exactly equal
// in contents and order.
func Equal2[K, V comparable](seq1 iter.Seq2[K, V], seq2 iter.Seq2[K, V]) bool {
	next1, stop1 := iter.Pull2(seq1)
	defer stop1()

	next2, stop2 := iter.Pull2(seq2)
	defer stop2()

	for {
		key1, value1, ok1 := next1()
		key2, value2, ok2 := next2()

		if !ok1 && !ok2 {
			break
		}

		if ok1 != ok2 {
			return false
		}

		if key1 != key2 {
			return false
		}

		if value1 != value2 {
			return false
		}
	}

	return true
}
