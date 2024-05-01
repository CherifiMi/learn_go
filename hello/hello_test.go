package main

import "testing"

func TestHello(t *testing.T) {
	g := hello()
	w := "hello"

	if g != w {
		t.Errorf("got %q wabts %q", g, w)
	}
}
