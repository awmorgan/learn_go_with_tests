package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {

		wallet := Wallet{balance: 10}
		wallet.Deposit(10)

		assertBalance(wallet, 20, t)
	})

	t.Run("Withdraw", func(t *testing.T) {

		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))

		assertBalance(wallet, 10, t)
		assertErr(err, nil, t)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBal := Bitcoin(20)
		wallet := Wallet{balance: startingBal}
		err := wallet.Withdraw(100)

		assertBalance(wallet, startingBal, t)
		assertErr(err, ErrInsufficientFunds, t)
	})

}

func assertErr(got error, want error, t *testing.T) {
	t.Helper()
	if got == nil && want != nil {
		t.Fatal("wanted an error but didn't get one")
	}
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertBalance(wallet Wallet, want Bitcoin, t testing.TB) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestBitcoin_String(t *testing.T) {
	tests := []struct {
		name string
		b    Bitcoin
		want string
	}{
		{name: "check string", b: 10, want: "10 BTC"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.String(); got != tt.want {
				t.Errorf("Bitcoin.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
