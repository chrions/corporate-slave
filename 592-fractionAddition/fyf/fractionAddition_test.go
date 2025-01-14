package fyf

import "testing"

func TestFraction(t *testing.T) {
	bs := []struct {
		Input    string
		Expected string
	}{
		{"-1/2+1/2", "0/1"},
		{"-1/2+1/2+1/3", "1/3"},
	}
	for _, b := range bs {
		output := fractionAddition(b.Input)
		if output != b.Expected {
			t.Errorf("input=>%s,expected=>%s,got=>%s", b.Input, b.Expected, output)
		}
	}
}
