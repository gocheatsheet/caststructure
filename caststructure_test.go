package caststructure

import (
	"testing"
)

func TestInterface(t *testing.T) {
	var from impl

	value, err := Interface(from, (*testA)(nil), (*TestB)(nil))
	if err != nil {
		t.Fatal(err)
	}

	if _, ok := value.(testA); !ok {
		t.Fatal("should implement A")
	}
	if _, ok := value.(TestB); !ok {
		t.Fatal("should implement B")
	}
	if _, ok := value.(testC); ok {
		t.Fatal("should not implement C")
	}
}

func TestInterface_nonImpl(t *testing.T) {
	from := 42
	_, err := Interface(from, (*testA)(nil))
	if err == nil {
		t.Fatal("should error")
	}
}

func TestInterface_nonPtr(t *testing.T) {
	var from impl
	_, err := Interface(from, (testA)(nil))
	if err == nil {
		t.Fatal("should error")
	}
}

type testA interface{ A() int }
type TestB interface{ B() int } // Purposefully exported to test that case
type testC interface{ C() int }

type impl struct{}

func (impl) A() int { return 42 }
func (impl) B() int { return 42 }
func (impl) C() int { return 42 }
