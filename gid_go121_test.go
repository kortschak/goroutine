// Copyright ©2023 Dan Kortschak. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build go1.21
// +build go1.21

package goroutine

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"testing"
)

func TestParentID(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 1000000; i++ {
		i := i
		wg.Add(2)
		go func() {
			defer wg.Done()
			parentID := ID()
			go func() {
				defer wg.Done()
				got := ParentID()
				want := parentGoid()
				if got != want {
					t.Errorf("unexpected parent id for goroutine number %d: mismatch with stack trace: got:%d want:%d", i, got, want)
				}
				if got != parentID {
					t.Errorf("unexpected parent id for goroutine number %d: mismatch with parent: got:%d want:%d", i, got, want)
				}
			}()
		}()
	}
	wg.Wait()
}

// parentGoid returns the parent goroutine ID extracted from a stack trace.
func parentGoid() int64 {
	var buf [1 << 10]byte
	n := runtime.Stack(buf[:], false)
	_, after, ok := strings.Cut(string(buf[:n]), " in goroutine ")
	if !ok {
		panic("cannot get parent goroutine id: no mark")
	}
	idField := strings.Fields(after)[0]
	id, err := strconv.ParseInt(idField, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("cannot get parent goroutine id: %v", err))
	}
	return id
}
