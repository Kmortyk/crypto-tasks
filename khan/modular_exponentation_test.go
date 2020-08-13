package khan

import "testing"

func Test_fastModExp(t *testing.T) {
	x, y, z := 98765, 1234, 123557

	result := fastModExp(x, y, z)
	if result != 70506 {
		t.Fatalf("wrong result: got=%v, want=%v", result, 70506)
	}
}
