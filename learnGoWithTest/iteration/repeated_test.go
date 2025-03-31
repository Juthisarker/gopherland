package iteration

import "testing"

func TestRepeat(t *testing.T){
	repeated := Repeated("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)

	}
}