package stack

import (
	"testing"

	"github.com/schbm/gotastructs/v2/general"
)

func TestStack(t *testing.T) {

	stringTests := []struct {
		Name   string
		List   general.Stack[rune]
		Data   string
		Result string
	}{
		{
			Name:   "array",
			List:   NewLinkedListStack[rune](),
			Data:   "HALLO",
			Result: "OLLAH",
		},
	}

	for _, test := range stringTests {
		for _, char := range test.Data {
			test.List.Push(char)
		}
		result := ""
		for test.List.Size() >= 1 {
			char, err := test.List.Pop()
			if err != nil {
				t.Error(err)
				t.Fail()
			}
			result += string(char)
		}
		if test.Result != result {
			t.Error("results do not match", test.Result, result)
			t.Fail()
		}
	}
}
