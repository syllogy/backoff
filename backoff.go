// Copyright © 2015-2020, Christian R. Vozar

// Package backoff provides stateless backoff policies for reconnect loops used
// to maintain persistent connections. Randomized wait times are utilized to
// avoid the thundering herd problem.
package backoff

import (
	"math/rand"
	"time"
)

func init() {
	// Seed random number generator
	rand.Seed(time.Now().UTC().UnixNano())
}

// Backoff interface is implemented by backoff policies.
type Backoff interface {
	Duration(n int) time.Duration
}

type RandomBackoff struct{}

// Duration returns calculated delays.
func (b RandomBackoff) Duration(n int) time.Duration {
	return time.Duration(rand.Int63())
}

// Policy implements Backoff, randomizing delays and saturating final
// value in Millis.
type Policy struct {
	Millis []int
}

// Default is a backoff policy ranging up to 5 seconds.
var Default = Policy{
	[]int{0, 10, 10, 100, 100, 500, 500, 3000, 3000, 5000},
}

// Duration returns the time duration of the n'th wait cycle in a
// backoff policy. This is b.Millis[n], randomized to avoid thundering herds.
func (b BackoffPolicy) Duration(n int) time.Duration {
	if n >= len(b.Millis) {
		n = len(b.Millis) - 1
	}

	return time.Duration(jitter(b.Millis[n])) * time.Millisecond
}

// jitter returns a random integer uniformly distributed in the range
// [0.5 * millis .. 1.5 * millis]
func jitter(millis int) int {
	if millis == 0 {
		return 0
	}

	return millis/2 + rand.Intn(millis)
}
