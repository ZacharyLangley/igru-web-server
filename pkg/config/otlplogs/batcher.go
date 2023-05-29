package otlplogs

import (
	"context"
	"time"
)

type batcher[T interface{}] struct {
	MaxItems int
	Timeout  time.Duration
}

func (b batcher[T]) Start(ctx context.Context, sync <-chan struct{}, in <-chan T) <-chan []T {
	output := make(chan []T, 10)
	go func() {
		defer close(output)
		timer := time.NewTimer(b.Timeout)
		buffer := make([]T, 0, b.MaxItems)
		for {
			select {
			case <-ctx.Done():
				// Flush buffer if it's not empty
				if len(buffer) > 0 {
					output <- buffer
				}
				return
			case <-timer.C:
				// Flush buffer if it's not empty and reset timer
				if len(buffer) > 0 {
					output <- buffer
					buffer = make([]T, 0, b.MaxItems)
				}
				timer.Reset(b.Timeout)
			case item, ok := <-in:
				if !ok {
					// Flush buffer if it's not empty
					if len(buffer) > 0 {
						output <- buffer
					}
					return
				}
				// Add item to buffer
				buffer = append(buffer, item)
				if len(buffer) >= b.MaxItems {
					// Flush buffer and reset timer
					output <- buffer
					buffer = make([]T, 0, b.MaxItems)
					timer.Reset(b.Timeout)
				}
			case <-sync:
				if len(buffer) > 0 {
					// Flush buffer and reset timer
					output <- buffer
					buffer = make([]T, 0, b.MaxItems)
					timer.Reset(b.Timeout)
				}
			}
		}
	}()
	return output
}
