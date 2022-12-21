package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	assert.Equal(t, 1747, Convert("1=-0-2"))
	assert.Equal(t, Convert("1=-0-2"), 1747)
	assert.Equal(t, Convert("12111"), 906)
	assert.Equal(t, Convert("2=0="), 198)
	assert.Equal(t, Convert("21"), 11)
	assert.Equal(t, Convert("2=01"), 201)
	assert.Equal(t, Convert("111"), 31)
	assert.Equal(t, Convert("20012"), 1257)
	assert.Equal(t, Convert("112"), 32)
	assert.Equal(t, Convert("1=-1="), 353)
	assert.Equal(t, Convert("1-12"), 107)
	assert.Equal(t, Convert("12"), 7)
	assert.Equal(t, Convert("1="), 3)
	assert.Equal(t, Convert("122"), 37)
}

func TestSolve1TestInput(t *testing.T) {
	testInput := `1=-0-2
12111
2=0=
21
2=01
111
20012
112
1=-1=
1-12
12
1=
122`

	result := Solve1(testInput)
	if result != "2=-1=0" {
		t.Errorf("got %v, want %v", result, "2=-1=0")
	}
}

func TestSolve1Input(t *testing.T) {
	input, _ := os.ReadFile("input.txt")

	result := Solve1(string(input))
	if result != "2-2=21=0021=-02-1=-0" {
		t.Errorf("got %v, want %v", result, "2-2=21=0021=-02-1=-0")
	}
}
