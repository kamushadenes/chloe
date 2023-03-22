package react

import (
	"testing"
)

func TestBytesWriter(t *testing.T) {
	writer := &BytesWriter{}
	expected := []byte("hello, world!")
	if _, err := writer.Write(expected); err != nil {
		t.Errorf("Unexpected error while writing bytes to writer: %v", err)
	}
	if err := writer.Close(); err != nil {
		t.Errorf("Unexpected error while closing writer: %v", err)
	}
	if len(writer.Bytes) != len(expected) {
		t.Errorf("BytesWriter did not write the correct number of bytes - expected %d, but got %d", len(expected), len(writer.Bytes))
	}
	for i, b := range writer.Bytes {
		if b != expected[i] {
			t.Errorf("BytesWriter wrote incorrect data - expected %v, but got %v", expected[i], b)
		}
	}
}
