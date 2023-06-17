// Copyright ©2023 Dan Kortschak. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build go1.21
// +build go1.21

package goroutine

import "unsafe"

// ParentID returns the runtime ID of goroutine that created the calling
// goroutine.
func ParentID() int64 {
	return idOf(getg(), parentGoidoff)
}

var parentGoidoff = offset("*runtime.g", "parentGoid")

// Link is a goroutine parent-child relationship.
type Link struct {
	Parent, Child int64
}

// All returns all the known goroutine parent-child relationships.
func All() []Link {
	var s []Link
	forEachG(func(g unsafe.Pointer) {
		s = append(s, Link{Parent: idOf(g, parentGoidoff), Child: idOf(g, goidoff)})
	})
	return s
}

//go:linkname forEachG runtime.forEachG
//go:nosplit
func forEachG(fn func(gp unsafe.Pointer))
