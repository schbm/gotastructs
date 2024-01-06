package set

import (
	"testing"

	"github.com/schbm/gotastructs/v1/element"
)

func TestHashSet(t *testing.T) {
	hs := NewHashSet()
	test := element.NewString("Bla")
	hs.Add(test)
	hs.Add(test)
	hs.Add(element.NewString("Bla"))
	hs.Add(element.NewString("Bla"))

	it := hs.Iterator()
	for it.HasNext() {
		t.Log(it.Next().String())
	}

	if hs.Size() != 3 { //t
		t.Error("duplicate entries")
	}
}
