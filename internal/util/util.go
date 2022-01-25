//
// Copyright Strimzi authors.
// License: Apache License 2.0 (see the file LICENSE or http://apache.org/licenses/LICENSE-2.0.html).
//

// Package util contains some utility functions
package util

import (
	"errors"
	"github.com/golang/glog"
	"io"
	"os"
	"syscall"
	"time"
)

//interface DeadLineExceededCounters {
//	Foo()
//}
// NowInMilliseconds returns the current time in milliseconds
func NowInMilliseconds() int64 {
	return time.Now().UnixNano() / 1000000
}

// IsDisconnection returns true if the err provided represents an error
func IsDisconnection(err error, deadlineExceededCount *int, deadlineExceededLimit int) bool {
	if errors.Is(err, os.ErrDeadlineExceeded) {
		*deadlineExceededCount++
		b := *deadlineExceededCount > deadlineExceededLimit
		if !b {
			glog.V(1).Infof("Ignoring %v as limit %d not reached yet (%d)", err, deadlineExceededLimit, *deadlineExceededCount)
		}
		return b
	} else {
		*deadlineExceededCount = 0
		return errors.Is(err, io.EOF) || errors.Is(err, syscall.ECONNRESET) || errors.Is(err, syscall.EPIPE) || errors.Is(err, syscall.ETIMEDOUT)
	}
}
