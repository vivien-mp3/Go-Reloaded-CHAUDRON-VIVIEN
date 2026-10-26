package main

import (
	//"errors"
	"go-reloaded/src"
	//"reflect"
	"testing"
)

func TestUp(t *testing.T) { //go test -v -run '^TestUp$'

	t.Run("Up", func(t *testing.T) {
		input := "They take the (up) bus to work. (up, 2)"
		want := "They take THE bus TO WORK."

		cuttedText := src.CutText(input)
		got := src.GlueText(cuttedText)


		if got != want {
			t.Errorf("%s Up%v = %v; want: %v", t.Name(), input, got, want)
			return
		}
	})
} 

func TestLow(t *testing.T) { //go test -v -run '^TestLow$'

	t.Run("Down", func(t *testing.T) {
		input := "SONIC. DEAD OR (low) ALIVE. IS MINE. (low, 2)"
		want := "SONIC. DEAD or ALIVE. is mine."

		cuttedText := src.CutText(input)
		got := src.GlueText(cuttedText)


		if got != want {
			t.Errorf("%s Down%v = %v; want: %v", t.Name(), input, got, want)
			return
		}
	})
} 

func TestCap(t *testing.T) { //go test -v -run '^TestCap$'

	t.Run("Cap", func(t *testing.T) {
		input := "The day of whom perished (cap, 2). Must be horrifying (cap)"
		want := "The day of Whom Perished. Must be Horrifying"

		cuttedText := src.CutText(input)
		got := src.GlueText(cuttedText)


		if got != want {
			t.Errorf("%s Down%v = %v; want: %v", t.Name(), input, got, want)
			return
		}
	})
} 

func TestGoReloaded(t *testing.T) { //go test -v -run '^TestGoReloaded$'

	t.Run("GoReloaded", func(t *testing.T) {
		input := "The quick BROWN (low) 'fox jumps   over the lazy dog (up, 2)  . pack my box with 5 (bin) dozen (cap) liquor jugs."
		want := "The quick brown 'fox' jumps over the LAZY DOG. Pack my box with 101 Dozen liquor jugs."

		cuttedText := src.CutText(input)
		got := src.GlueText(cuttedText)


		if got != want {
			t.Errorf("%s Down%v = %v; want: %v", t.Name(), input, got, want)
			return
		}
	})
} 