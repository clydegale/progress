// Copyright 2013 Jimmy Zelinskie. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package progress implements a thread-safe, ASCII progress bar for terminals.
package progress

import (
	"fmt"
	"io"
	"sync"
	"sync/atomic"
	"time"
)

const (
	termWidth     = 80
	barWidth      = termWidth - 7
	barBegin      = "["
	barEnd        = "]"
	barComplete   = "#"
	barIncomplete = "-"
)

type ProgressBar struct {
	mu  sync.Mutex
	out io.Writer

	max   int64
	count int64
}

// New returns a new progress bar that will be complete when it has been
// incremented max number of times.
func New(out io.Writer, max int64) *ProgressBar {
	return &ProgressBar{
		out: out,
		max: max,
	}
}

// Complete returns the current number of times Increment() has been called.
func (pb ProgressBar) Complete() int64 {
	return pb.count
}

// Incomplete returns the current number of times Increment() needs to be
// called to reach completion.
func (pb ProgressBar) Incomplete() int64 {
	return pb.max - pb.count
}

// Completed determines whether the progress has reached completion.
func (pb ProgressBar) Completed() bool {
	return pb.count >= pb.max
}

// CurrentCompletion returns the current ratio of completion.
func (pb *ProgressBar) CurrentCompletion() float64 {
	return float64(pb.count) / float64(pb.max)
}

// CurrentPercentage returns the current percentage of completion.
func (pb *ProgressBar) CurrentPercentage() int {
	return int(pb.CurrentCompletion() * 100)
}

// Increment atomically increments the progress of the bar by 1.
func (pb *ProgressBar) Increment() {
	atomic.AddInt64(&pb.count, 1)
}

// IncrementBy atomically increments the progress of the bar by a provided value.
func (pb *ProgressBar) IncrementBy(x int64) {
	atomic.AddInt64(&pb.count, x)
}

// Draw() atomically writes the bar to the provided io.Writer.
func (pb *ProgressBar) Draw() (err error) {
	pb.mu.Lock()
	defer pb.mu.Unlock()

	ratio := pb.CurrentCompletion()
	currentPercent := int(ratio * 100)
	percentStr := fmt.Sprintf(" %d%%", currentPercent)

	// Build the bar
	complete := ratio * barWidth
	incomplete := barWidth - complete
	bar := barBegin
	for i := float64(0); i < complete; i++ {
		bar += barComplete
	}
	for i := float64(0); i < incomplete; i++ {
		bar += barIncomplete
	}
	bar += barEnd

	_, err = pb.out.Write([]byte("\r" + bar + percentStr))
	return
}

// DrawEvery spawns a new goroutine that writes the bar to the provided io.Writer
// at the given duration until the bar has completed.
func (pb *ProgressBar) DrawEvery(t time.Duration) {
	go func() {
		pb.Draw()
		for !pb.Completed() {
			time.Sleep(t)
			pb.Draw()
		}
	}()
}
