// Copyright 2013 Jimmy Zelinskie. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package progress

import (
	"bytes"
	"testing"
)

func TestComplete(t *testing.T) {
	var buf bytes.Buffer
	pb := New(&buf, 100)

	for i := 0; i < 100; i++ {
		pb.Increment()
	}
	pb.Draw()

	golden := []byte("\r[##############################################################################] 100%")
	if !bytes.Equal(golden, buf.Bytes()) {
		t.Fatalf("got: %s wanted: %s", buf.Bytes(), golden)
	}
}

func TestLeft(t *testing.T) {
	var buf bytes.Buffer
	pb := New(&buf, 100)

	for i := 1; i <= 100; i++ {
		pb.Increment()
		if pb.Left() != int64(100-i) {
			t.Fatalf("got: %d wanted: %d", pb.Left(), 100-i)
		}
	}
}

func TestProgress(t *testing.T) {
	var buf bytes.Buffer
	pb := New(&buf, 100)

	for i := 1; i <= 100; i++ {
		pb.Increment()
		if pb.Progress() != int64(i) {
			t.Fatalf("got: %d wanted: %d", pb.Progress(), 100-i)
		}
	}
}
