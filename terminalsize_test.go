package main

import "testing"

func TestConsoleSize(t *testing.T) {
	expectedw, expectedh := 58, 118 // stty size static size!!!
	w, h := ConsoleSize()
	if expectedw == w && expectedh == h {
		t.Errorf("Failed ! got %v %v want %v %v", w, h, expectedw, expectedh)
	} else {
		t.Logf("Success !")
	}

}
