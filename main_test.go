package main

import "testing"

func TestSum(t *testing.T) {
    got := Sum(5, 3)
    want := 8
    
    if got != want {
        t.Errorf("Sum(5, 3) = %d; want %d", got, want)
    }
}
