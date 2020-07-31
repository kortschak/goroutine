// Copyright ©2020 Dan Kortschak. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goroutine

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"testing"
)

func TestID(t *testing.T) {
	got := ID()
	want := goid()
	if got != want {
		t.Fatalf("unexpected id for main goroutine: got:%d want:%d", got, want)
	}
	var wg sync.WaitGroup
	for i := 0; i < 100000; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			got := ID()
			want := goid()
			if got != want {
				t.Errorf("unexpected id for goroutine number %d: got:%d want:%d", i, got, want)
			}
		}()
	}
	wg.Wait()
}

// goid returns the goroutine ID extracted from a stack trace.
func goid() int64 {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.ParseInt(idField, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}
