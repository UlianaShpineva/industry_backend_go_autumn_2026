package main

import (
	"context"
	"errors"
	"sync"
)

type Result[R any] struct {
	Value R
	Err   error
}

var ErrInvalidWorkers = errors.New("workers must be positive")

func ParallelMap[T any, R any](
	ctx context.Context,
	workers int,
	in []T,
	fn func(context.Context, T) (R, error),
) ([]Result[R], error) {
	if workers <= 0 {
		return nil, ErrInvalidWorkers
	}
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	results := make([]Result[R], len(in))
	if len(in) == 0 {
		return results, nil
	}

	w := workers
	if w > len(in) {
		w = len(in)
	}

	jobs := make(chan int, len(in))
	for i := range in {
		jobs <- i
	}
	close(jobs)

	var wg sync.WaitGroup
	wg.Add(w)

	for i := 0; i < w; i++ {
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case idx, ok := <-jobs:
					if !ok {
						return
					}
					if ctx.Err() != nil {
						return
					}

					v, err := fn(ctx, in[idx])
					if err != nil {
						results[idx] = Result[R]{Err: err}
					} else {
						results[idx] = Result[R]{Value: v}
					}
				}
			}
		}()
	}

	wg.Wait()

	if err := ctx.Err(); err != nil {
		return nil, err
	}
	return results, nil
}
