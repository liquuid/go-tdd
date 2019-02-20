package pe

import "testing"

func TestWallet(t *testing.T){
	assertBalance := func(t *testing.T, wallet Wallet,want Bitcoin){
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s, want %s", got , want )
		}
	}
	assertError := func(t *testing.T, got error, want error){
		t.Helper()
		if got == nil {
			t.Fatal("wanted an error but didnt get one")
		}
		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	}

	assertNoError := func(t *testing.T, got error) {
		t.Helper()
		if got != nil {
			t.Fatal("got an error but didnt want one")
		}
	}

	t.Run("put 10", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, 10)
	})

	t.Run("put 20", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(20))

		assertBalance(t, wallet, 20)

	})

	t.Run("put 20, withdraw 5", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(20))
		err := wallet.Withdraw(Bitcoin(5))
		assertBalance(t, wallet, 15)
		assertNoError(t, err)
	})
	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficientFunds)
	})

}