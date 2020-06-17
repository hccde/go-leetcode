package main

import (
	"testing"
)

func TestisMatch(t *testing.T){
	r := isMatch("aaa","ab*a*c*a");
	if r != false {
		t.Errorf("error case: one")
	}
}