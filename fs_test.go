// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gofs_test

import (
	"fmt"
	"testing"

	gofs "github.com/yungao/go-http-fs"
)

func TestServeFile(t *testing.T) {
	gofs.ServeFile(nil, nil, "testdata/file", func(start, send int64) {
		fmt.Printf("progress downloading, start: %d, send: %d \n", start, send)
	})
}
