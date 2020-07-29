package main

import ("testing")

func TestTenFibonacci (t *testing.T) {
	sequence := fibonacci(10)
	if len(sequence) != 7 {
		t.Fail()
	}

	if (sequence[0] != 0 && sequence[1] != 1 && sequence[2] != 1 && sequence[3] != 2 && sequence[4] != 3 && sequence[5] != 5 && sequence [6] != 8) {
		t.Fail()
	}
}

func TestZeroFibonacci (t *testing.T) {
	sequence := fibonacci(0)
	if len(sequence) != 1 {
		t.Fail()
	}

	if (sequence[0] != 0) {
		t.Fail()
	}
}

func TestOneFibonacci (t *testing.T) {
	sequence := fibonacci(1)
	if len(sequence) != 2 {
		t.Fail()
	}

	if (sequence[0] != 0 && sequence[1] != 1) {
		t.Fail()
	}
}
