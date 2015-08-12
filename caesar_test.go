package main

import "testing"

func TestAShift1ShouldBeB(t *testing.T) {
	c := Caesar("A", 1)

	if c != "B" {
		t.Errorf(`expect "B" but got %s`, c)
	}
}

func TestAShift2ShouldBeC(t *testing.T) {
	c := Caesar("A", 2)

	if c != "C" {
		t.Errorf(`expect "C" but got %s`, c)
	}
}

func TestAShift3ShouldBeD(t *testing.T) {
	c := Caesar("A", 3)

	if c != "D" {
		t.Errorf(`expect "D" but got %s`, c)
	}
}

func TestAShift3ShouldBeE(t *testing.T) {
	c := Caesar("A", 4)

	if c != "E" {
		t.Errorf(`expect "E" but got %s`, c)
	}
}
