// Copyright ©2023 Dan Kortschak. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build go1.21 && !go1.23
// +build go1.21,!go1.23

package goroutine

import "unsafe"

// Link is a goroutine parent-child relationship.
//
// Only available for go1.21 and go1.22.
type Link struct {
	Parent, Child int64
}

// All returns all the known goroutine parent-child relationships.
//
// Only available for go1.21 and go1.22.
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
