package main

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(a, b string) (string, error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(10 * time.Second):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(b string) chan struct{} {
	var ch = make(chan struct{})
	go func() {
		http.Get(b)
		close(ch)
	}()
	return ch
}
