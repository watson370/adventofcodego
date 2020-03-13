package graph

import "testing"

func TestIntcode(t *testing.T) {
	mycicleb := RingBuffer{
		buffer: make([]int, 10),
		in:     0,
		out:    0,
	}
	if mycicleb.Full() {
		t.Error("empty buffer returned full")
	}
	if !mycicleb.Empty() {
		t.Error("empty buffer didn't return empty")
	}
	for i := 0; i < 8; i++ {
		mycicleb.Add(i)
	}
	if mycicleb.Full() {
		t.Error("almost full buffer returned full")
	}
	if mycicleb.Empty() {
		t.Error("partial buffer returned empty")
	}
	mycicleb.Add(8)
	if !mycicleb.Full() {
		t.Error("full buffer returned not full")
	}
	if mycicleb.Empty() {
		t.Error("full buffer says empty")
	}
	if val, ok := mycicleb.Peek(); val != 0 && ok {
		t.Error("peek should be 0")
	}
	for i := 0; i < 8; i++ {
		mycicleb.Pop()
	}
	if mycicleb.Empty() {
		t.Error("partial buffer returned empty")
	}
	mycicleb.Pop()
	if !mycicleb.Empty() {
		t.Error("empty buffer didn't return empty")
	}
	val, ok := mycicleb.Pop()
	if val != 0 || ok {
		t.Error("pop from empty returned ", val, ", ok val of ", ok)
	}

}
