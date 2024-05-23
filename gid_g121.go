// Copyright ©2023 Dan Kortschak. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build go1.21
// +build go1.21

package goroutine

// ParentID returns the runtime ID of goroutine that created the calling
// goroutine.
func ParentID() int64 {
	return idOf(getg(), parentGoidoff)
}

var parentGoidoff = offset("*runtime.g", "parentGoid")
