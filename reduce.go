package gloop

import "iter"

// ReduceFunc is the function signature of the reducing function used
// in [Reduce].
type ReduceFunc[V any] func(V, V) V

// Reduce runs a given function on each adjacent pair in an iter.Seq
// sequence and accumulates the result into a single value.
func Reduce[V any](seq iter.Seq[V], f ReduceFunc[V]) V {
	var reducedValue V
	first := true

	for value := range seq {
		if first {
			reducedValue = value
			first = false
			continue
		}

		reducedValue = f(reducedValue, value)
	}

	return reducedValue
}

// Reduce2Func is the function signature of the reducing function used
// in [Reduce2].
type Reduce2Func[K, V any] func(K, V, K, V) (K, V)

// Reduce2 runs a given function on each adjacent pair of keys and
// values in an iter.Seq2 sequence and accumulates the result into a
// single key and value pair.
func Reduce2[K, V any](seq iter.Seq2[K, V], f Reduce2Func[K, V]) (K, V) {
	var reducedKey K
	var reducedValue V
	first := true

	for key, value := range seq {
		if first {
			reducedKey = key
			reducedValue = value
			first = false
			continue
		}

		reducedKey, reducedValue = f(reducedKey, reducedValue, key, value)
	}

	return reducedKey, reducedValue
}
