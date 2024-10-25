package main

import (
	"testing"
	"time"
)

func Test_dine(t *testing.T) {
	eatTime = 0 * time.Second
	thinkTime = 0 * time.Second

	for i := 0; i < 10; i++ {
		orderFinished = []string{}
		dine()
		if len(orderFinished) != len(philosophers) {
			t.Errorf("Incorrect length of slice, expected %v, got %v", len(philosophers), len(orderFinished))

		}
	}
}

func Test_dineWithVaryingDelays(t *testing.T) {
	var tests = []struct {
		name  string
		delay time.Duration
	}{
		{"zero delay", 0 * time.Second},
		{"quater second delay", 0 * time.Millisecond * 250},
		{"half second delay", 0 * time.Millisecond * 500},
	}

	for _, e := range tests {
		orderFinished = []string{}

		eatTime = e.delay
		thinkTime = e.delay

		dine()
		if len(orderFinished) != len(philosophers) {
			t.Errorf("Incorrect length of slice, expected %v, got %v", len(philosophers), len(orderFinished))
		}
	}

}
