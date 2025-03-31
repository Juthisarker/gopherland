package main

import "testing"

func TestHello(t *testing.T) {
// 	t.Run("saying hello to people", func(t *testing.T) {
	
// 	got := Hello("juthi")
// 	want := "Hello, juthi"

//     assertCorrectMessage(t, got, want)
// })
// t.Run("saying hello when empty string is supplied", func(t *testing.T) {
// 	got := Hello("")
// 	want := "Hello, World"

//     assertCorrectMessage(t, got, want) 
// })
t.Run("in spanish", func(t *testing.T) {
	got := Hello("Elodie", "Spanish")
	want := "Hola, Elodie"
	
	assertCorrectMessage(t, got, want)
})
}

func assertCorrectMessage(t testing.TB, got, want string){
  //   t.helper()
	 if got != want {
		t.Errorf("got %q want %q", got, want)
	 }
}