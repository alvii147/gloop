package gloop

import (
	"context"
	"iter"
	"sync"
)

// ParallelizeOptions defines configurable options for Parallelize.
type ParallelizeOptions struct {
	// Context is used to send a cancel signal.
	Context context.Context
	// MaxThreads defines the maximum number of concurrent threads
	// allowed. If nil, there is no maximum.
	MaxThreads *int
}

// ParallelizeOptionFunc is the function signature of configuration
// helpers for Parallelize.
type ParallelizeOptionFunc func(*ParallelizeOptions)

// WithParallelizeContext is a helper for configuring context in
// Parallelize.
func WithParallelizeContext(ctx context.Context) ParallelizeOptionFunc {
	return func(o *ParallelizeOptions) {
		o.Context = ctx
	}
}

// WithParallelizeMaxThreads is a helper for configuring maximum number
// of concurrent threads in Parallelize.
func WithParallelizeMaxThreads(maxThreads int) ParallelizeOptionFunc {
	return func(o *ParallelizeOptions) {
		o.MaxThreads = &maxThreads
	}
}

// ParallelizeFunc is the function signature of the function to be
// parallelized in Parallelize.
type ParallelizeFunc[V any] func(V)

// Parallelize runs a function on each value in an iter.Seq sequence on
// separate goroutines.
func Parallelize[V any](
	seq iter.Seq[V],
	f ParallelizeFunc[V],
	opts ...ParallelizeOptionFunc,
) {
	Parallelize2(Enumerate(seq), func(_ int, value V) {
		f(value)
	}, opts...)
}

// Parallelize2Func is the function signature of the function to be
// parallelized in Parallelize2.
type Parallelize2Func[K, V any] func(K, V)

// Parallelize2 runs a function on each value in an iter.Seq2 sequence
// on separate goroutines.
func Parallelize2[K, V any](
	seq iter.Seq2[K, V],
	f Parallelize2Func[K, V],
	opts ...ParallelizeOptionFunc,
) {
	options := ParallelizeOptions{
		Context:    context.Background(),
		MaxThreads: nil,
	}

	for _, opt := range opts {
		opt(&options)
	}

	ctx := context.Background()
	if options.Context != nil {
		ctx = options.Context
	}

	var semaphore chan struct{}
	if options.MaxThreads != nil {
		semaphore = make(chan struct{}, *options.MaxThreads)
		defer close(semaphore)
	}

	var wg sync.WaitGroup

	for key, value := range seq {
		if semaphore != nil {
			semaphore <- struct{}{}
		}

		wg.Add(1)
		go func(k K, v V) {
			defer wg.Done()
			if semaphore != nil {
				defer func(s <-chan struct{}) {
					<-s
				}(semaphore)
			}

			select {
			case <-ctx.Done():
				return
			default:
				f(k, v)
				return
			}
		}(key, value)
	}

	wg.Wait()
}
