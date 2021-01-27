package hello

import (
	"testing"
)

func TestHello(t *testing.T) {
	hello := Hello()
	if hello != "Hello!" {
		t.Fatalf("Hello() returned %s", hello)
	}
}
