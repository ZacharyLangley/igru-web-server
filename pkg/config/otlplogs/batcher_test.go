package otlplogs

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBatcher_FillUp(t *testing.T) {
	t.Parallel()
	batch := batcher[int]{
		MaxItems: 10,
		Timeout:  10 * time.Second,
	}
	input := make(chan int, 10)
	ctx, done := context.WithCancel(context.Background())
	syncChan := make(chan struct{})
	output := batch.Start(ctx, syncChan, input)
	// Add 10 integers
	want := make([]int, 10)
	for i := 0; i < 10; i++ {
		want[i] = i
	}
	for _, val := range want {
		input <- val
	}
	// Expect output to be a list of all 10 numbers
	got := <-output
	assert.Equal(t, want, got)
	// Cancel context and observe that the output is closed
	done()
	val, ok := <-output
	assert.False(t, ok, "value", val)
}

func TestBatcher_ManualSync(t *testing.T) {
	t.Parallel()
	batch := batcher[int]{
		MaxItems: 10,
		Timeout:  10 * time.Second,
	}
	input := make(chan int)
	ctx, done := context.WithCancel(context.Background())
	syncChan := make(chan struct{})
	output := batch.Start(ctx, syncChan, input)
	// Add 5 integers
	want := make([]int, 5)
	for i := 0; i < 5; i++ {
		want[i] = i
	}
	for _, val := range want {
		input <- val
	}
	// Send sync channel
	syncChan <- struct{}{}
	// Expect output to be a list of all 10 numbers
	got := <-output
	assert.Equal(t, want, got)
	// Cancel context and observe that the output is closed
	done()
	val, ok := <-output
	assert.False(t, ok, "value", val)
}

func TestBatcher_Timeout(t *testing.T) {
	t.Parallel()
	batch := batcher[int]{
		MaxItems: 10,
		Timeout:  1 * time.Second,
	}
	input := make(chan int)
	ctx, done := context.WithCancel(context.Background())
	syncChan := make(chan struct{})
	output := batch.Start(ctx, syncChan, input)
	// Add 5 integers
	want := make([]int, 5)
	for i := 0; i < 5; i++ {
		want[i] = i
	}
	for _, val := range want {
		input <- val
	}
	// wait for timeout
	// Expect output to be a list of all 10 numbers
	got := <-output
	assert.Equal(t, want, got)
	// Cancel context and observe that the output is closed
	done()
	val, ok := <-output
	assert.False(t, ok, "value", val)
}

func TestBatcher_CloseInput(t *testing.T) {
	t.Parallel()
	batch := batcher[int]{
		MaxItems: 10,
		Timeout:  1 * time.Second,
	}
	input := make(chan int)
	ctx := context.Background()
	syncChan := make(chan struct{})
	output := batch.Start(ctx, syncChan, input)
	// Add 5 integers
	want := make([]int, 5)
	for i := 0; i < 5; i++ {
		want[i] = i
	}
	for _, val := range want {
		input <- val
	}
	// close input for timeout
	close(input)
	// Expect output to be a list of all 10 numbers
	got := <-output
	assert.Equal(t, want, got)
	// Cancel context and observe that the output is closed
	val, ok := <-output
	assert.False(t, ok, "value", val)
}

func TestBatcher_FlushOnExit(t *testing.T) {
	t.Parallel()
	batch := batcher[int]{
		MaxItems: 10,
		Timeout:  1 * time.Second,
	}
	input := make(chan int)
	ctx, done := context.WithCancel(context.Background())
	syncChan := make(chan struct{})
	output := batch.Start(ctx, syncChan, input)
	// Add 5 integers
	want := make([]int, 5)
	for i := 0; i < 5; i++ {
		want[i] = i
	}
	for _, val := range want {
		input <- val
	}
	// close context
	done()
	// Expect output to be a list of all 10 numbers
	got := <-output
	assert.Equal(t, want, got)
	val, ok := <-output
	assert.False(t, ok, "value", val)
}
