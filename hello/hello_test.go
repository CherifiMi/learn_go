package main

import "testing"

func TestHello(t *testing.T) {
	g := hello("mito")
	w := "hello, mito"

	if g != w {
		t.Errorf("got %q wabts %q", g, w)
	}
}
