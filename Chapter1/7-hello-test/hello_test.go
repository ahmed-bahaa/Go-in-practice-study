package main

import "testing"

func TestName(t *testing.T) {

	name := getName()

	if name != "world" {
		t.Error("Response from getNmae func is unexpected value")
	}
}
