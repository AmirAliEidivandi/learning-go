package main

import "testing"

func TestIsvalidEmail(t *testing.T) {
	data := "email@example.com"
	if !isValidEmail(data) {
		t.Errorf("IsValidEmail(%v)=false, want true", data)
	}
}