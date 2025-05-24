package gloop_test

import (
	"context"
	"sync/atomic"
	"testing"
	"time"

	"github.com/alvii147/gloop"
	"github.com/stretchr/testify/require"
)

func TestWithParallelizeContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	options := gloop.ParallelizeOptions{}
	gloop.WithParallelizeContext(ctx)(&options)

	ctxDone := false
	select {
	case <-options.Context.Done():
		ctxDone = true
	default:
	}
	require.False(t, ctxDone)

	cancel()

	ctxDone = true
	select {
	case <-options.Context.Done():
	default:
		ctxDone = false
	}
	require.True(t, ctxDone)
}

func TestWithParallelizeMaxThreads(t *testing.T) {
	maxThreads := 42
	options := gloop.ParallelizeOptions{}
	gloop.WithParallelizeMaxThreads(maxThreads)(&options)

	require.NotNil(t, options.MaxThreads)
	require.Equal(t, maxThreads, *options.MaxThreads)
}

func TestParallelizeSlice(t *testing.T) {
	values := []string{"a", "b", "c"}
	valuesCh := make(chan string, len(values))

	done := make(chan struct{}, 1)
	channelOverflow := false

	go func() {
		gloop.Parallelize(gloop.Slice(values), func(v string) {
			select {
			case valuesCh <- v:
			default:
				channelOverflow = true
			}
		})
		done <- struct{}{}
	}()

	select {
	case <-done:
	case <-time.After(time.Second * 10):
		t.Fatal("done signal took too long")
	}

	require.False(t, channelOverflow)

	close(valuesCh)

	gotValues := make([]string, 0)

	for v := range valuesCh {
		gotValues = append(gotValues, v)
	}

	require.ElementsMatch(t, values, gotValues)
}

func TestParallelizeString(t *testing.T) {
	s := "FizzBuzz"
	rCh := make(chan rune, len(s))

	done := make(chan struct{}, 1)
	channelOverflow := false

	go func() {
		gloop.Parallelize(gloop.String(s), func(r rune) {
			select {
			case rCh <- r:
			default:
				channelOverflow = true
			}
		})
		done <- struct{}{}
	}()

	select {
	case <-done:
	case <-time.After(time.Second * 10):
		t.Fatal("done signal took too long")
	}

	require.False(t, channelOverflow)

	close(rCh)

	gotRunes := make([]rune, 0)

	for r := range rCh {
		gotRunes = append(gotRunes, r)
	}

	require.ElementsMatch(t, []rune(s), gotRunes)
}

func TestParallelize2(t *testing.T) {
	m := map[string]int{
		"Fizz": 3,
		"Buzz": 1,
		"Bazz": 4,
	}
	keysCh := make(chan string, len(m))
	valuesCh := make(chan int, len(m))

	done := make(chan struct{}, 1)
	channelOverflow := false

	go func() {
		gloop.Parallelize2(gloop.Map(m), func(k string, v int) {
			select {
			case keysCh <- k:
			default:
				channelOverflow = true
			}

			select {
			case valuesCh <- v:
			default:
				channelOverflow = true
			}
		})
		done <- struct{}{}
	}()

	select {
	case <-done:
	case <-time.After(time.Second * 10):
		t.Fatal("done signal took too long")
	}

	require.False(t, channelOverflow)

	close(keysCh)
	close(valuesCh)

	gotKeys := make([]string, 0)
	for k := range keysCh {
		gotKeys = append(gotKeys, k)
	}

	gotValues := make([]int, 0)
	for v := range valuesCh {
		gotValues = append(gotValues, v)
	}

	for k, v := range m {
		require.Contains(t, gotKeys, k)
		require.Contains(t, gotValues, v)
	}
}

func TestParallelizeCancelContext(t *testing.T) {
	values := []string{"Fizz"}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	done := make(chan struct{}, 1)
	functionCalled := false

	go func() {
		gloop.Parallelize(gloop.Slice(values), func(v string) {
			functionCalled = true
		}, gloop.WithParallelizeContext(ctx))
		done <- struct{}{}
	}()

	select {
	case <-done:
	case <-time.After(time.Second * 10):
		t.Fatal("done signal took too long")
	}

	require.False(t, functionCalled)
}

func TestParallelize2CancelContext(t *testing.T) {
	values := []string{"Fizz"}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	done := make(chan struct{}, 1)
	functionCalled := false

	go func() {
		gloop.Parallelize2(gloop.Enumerate(gloop.Slice(values)), func(_ int, _ string) {
			functionCalled = true
		}, gloop.WithParallelizeContext(ctx))
		done <- struct{}{}
	}()

	select {
	case <-done:
	case <-time.After(time.Second * 10):
		t.Fatal("done signal took too long")
	}

	require.False(t, functionCalled)
}

func TestParallelizeSingleThreaded(t *testing.T) {
	values := []string{"a", "b", "c"}

	var (
		concurrentCallers    atomic.Int64
		maxConcurrentCallers atomic.Int64
	)

	done := make(chan struct{}, 1)
	go func() {
		gloop.Parallelize(gloop.Slice(values), func(_ string) {
			concurrentCallers.Add(1)
			defer concurrentCallers.Add(-1)

			maxConcurrentCallers.Store(max(maxConcurrentCallers.Load(), concurrentCallers.Load()))
		}, gloop.WithParallelizeMaxThreads(1))
		done <- struct{}{}
	}()

	select {
	case <-done:
	case <-time.After(time.Second * 10):
		t.Fatal("done signal took too long")
	}

	require.EqualValues(t, 1, maxConcurrentCallers.Load())
}

func TestParallelize2SingleThreaded(t *testing.T) {
	values := []string{"a", "b", "c"}

	var (
		concurrentCallers    atomic.Int64
		maxConcurrentCallers atomic.Int64
	)

	done := make(chan struct{}, 1)
	go func() {
		gloop.Parallelize2(gloop.Enumerate(gloop.Slice(values)), func(_ int, _ string) {
			concurrentCallers.Add(1)
			defer concurrentCallers.Add(-1)

			maxConcurrentCallers.Store(max(maxConcurrentCallers.Load(), concurrentCallers.Load()))
		}, gloop.WithParallelizeMaxThreads(1))
		done <- struct{}{}
	}()

	select {
	case <-done:
	case <-time.After(time.Second * 10):
		t.Fatal("done signal took too long")
	}

	require.EqualValues(t, 1, maxConcurrentCallers.Load())
}
