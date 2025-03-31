package pointersanderrors
import "fmt"
import "testing"
// TestWallet is a test function for the Wallet struct
func TestWallet(t *testing.T){
	wallet:= Wallet{}
	wallet.Deposit(10)

	got := wallet.Balance()

	fmt.Printf("address of balance in test is %p \n", &wallet.balance)
	expected := 10

	if got != expected {
		t.Errorf("got %d want %d", got, expected)
	}
}