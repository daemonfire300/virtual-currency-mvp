package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestAssertions(t *testing.T) {
	user := NewUser("Frank", 1234)
	currency := NewCurrency("FruitCoins")
	wallet := NewWallet(user, 123100.0, currency)
	walletC := NewWallet(user, 123100.0, currency)

	currencyB := NewCurrency("FruitCoinsCS")
	walletB := NewWallet(user, 123100.0, currencyB)

	Convey("Relations tests", t, func() {
		So(nil, ShouldBeNil)

		So(wallet.Currency, ShouldEqual, currency)
		So(wallet.Owner, ShouldEqual, user)
		So(wallet.Volume, ShouldBeGreaterThan, 0.0)
		So(user.Name, ShouldEqual, "Frank")
	})

	Convey("Wallet tests", t, func() {
		So(wallet.Currency, ShouldEqual, currency)
		So(wallet.Owner, ShouldEqual, user)
		So(wallet.Volume, ShouldBeGreaterThan, 0.0)
		state, err := wallet.Subtract(123100.0)
		So(state, ShouldEqual, true)
		So(err, ShouldBeNil)
		So(wallet.Volume, ShouldEqual, 0.0)
	})

	Convey("Transaction Wrong Currencies", t, func() {
		transaction := NewTransaction(10000.0, wallet, walletB)
		state, err := transaction.Commit()
		So(state, ShouldEqual, false)
		So(err, ShouldNotBeNil)
	})

	Convey("Transaction Right Currencies", t, func() {
		transaction := NewTransaction(100.0, walletC, walletC)
		state, err := transaction.Commit()
		So(state, ShouldEqual, true)
		So(err, ShouldBeNil)
	})
}

func panics() {
	panic("Goofy Gophers!")
}
