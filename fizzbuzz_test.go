package main

import "testing"

func TestRun(t *testing.T) {
	// what? 
	t.Log("Just doing a run of it")
	derp := fbdetection(20)
	if derp != "Buzz" {
		t.Errorf("expected Buzz, did not get Buzz")
	}	
	t.Log("this is what the output is", derp)
}

func TestFuckBitches(t *testing.T) {
	t.Log("fuck bitches")
}