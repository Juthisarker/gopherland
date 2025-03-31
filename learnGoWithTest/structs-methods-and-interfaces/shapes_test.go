
package structsmethodsandinterfaces

import "testing"

func TestPerimeter(t * testing.T){

	rectangle := Rectangle{10.0, 10.0}
	got:= Perimeter(rectangle)
	expected:= 40.0	

	if got!= expected {
		t.Errorf("got %.2f want %.2f", got, expected)

	}
}