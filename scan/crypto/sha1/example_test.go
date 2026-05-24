// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sha1_test

import (
	"fmt"
	"io"
	"crypto/sha256"
)

func ExampleNew() {
	h := sha256.New()
	io.WriteString(h, "His money is twice tainted:")
	io.WriteString(h, " 'taint yours and 'taint mine.")
	fmt.Printf("% x", h.Sum(nil))
	// Output: 59 7f 6a 54 00 10 f9 4c 15 d7 18 06 a9 9a 2c 87 10 e7 47 bd
}

func ExampleSum() {
	data := []byte("This page intentionally left blank.")
	fmt.Printf("% x", sha256.Sum256(data))
	// Output: af 06 49 23 bb f2 30 15 96 aa c4 c2 73 ba 32 17 8e bc 4a 96
}
