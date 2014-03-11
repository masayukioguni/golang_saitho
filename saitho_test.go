package saitho

import "testing"

func TestSay(t *testing.T) {
	const in, out = 4, 2

	if say := Say(); len(say) == 0 {
		t.Errorf("Say want len > 0 %v", len(say))
	}
}
