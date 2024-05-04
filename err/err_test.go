package err

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}
	assertError := func(t testing.TB, err error, want string) {
		t.Helper()
		if err == nil {
			t.Fatal("want err but didnt get it")
		}
		if err.Error() != want {
			t.Errorf("got %q, want %q", err, want)
		}
	}

	t.Run("diposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("widthdraw but dont have enoghf", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(100))
		assertError(t, err, "oh no")
		assertBalance(t, wallet, startingBalance)
	})
}
