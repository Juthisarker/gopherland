package main
import "reflect" 
import "testing"

func TestSum(t *testing.T){

	t.Run("collection of 5 numbers", func(t *testing.T) {
	numbers := []int{1,2,3,4,5}

	got := Sum(numbers)
	expected:= 15

	if got != expected {
		t.Errorf("got %d want %d given, %v", got, expected, numbers)
	}
})

t.Run("collection of any numbers", func(t *testing.T) {

	numbers:= []int{1,2,3}
	got:= Sum(numbers)
	expected:= 6

	if got!= expected {
		t.Errorf("got %d want %d given, %v", got, expected, numbers)
	}
  })
}

func TestSumAll(t *testing.T){

	got:= SumAll([]int {1,2}, []int {0,9})
	want:= []int{3,9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}


func TestSumAllTrails(t *testing.T) {
	t.Run("make the sums of all slices", func(t *testing.T){
	got:= SumAllTrails([]int{1,2}, []int{0,9})
	want:= []int{2,9}

	if !reflect.DeepEqual(got, want){
		t.Errorf("got %v want %v", got, want)
	}

})

t.Run("safely sum empty slices", func(t* testing.T){
	got:= SumAllTrails([]int{}, []int{3,4,5})
	want:= []int{0,9}

	if !reflect.DeepEqual(got, want){
		t.Errorf("got %v want %v", got, want)
	}
})

}