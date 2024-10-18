package main

import "testing"

func Test_updateMessage(t *testing.T) {

	msg = "Hello, world!"

	wg.Add(2)
	go updateMessage("x")
	go updateMessage("This is a test")
	wg.Wait()

	if msg != "This is a test" {
		t.Error("Expected 'This is a test' but got", msg)
	}
}
