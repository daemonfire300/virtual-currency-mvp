package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestAssertions(t *testing.T) {
	user := NewUser("Frank", 1234)
	currency := NewCurrency("FruitCoins")
	wallet := NewWallet(user, 123100.01210121, currency)

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
		state, err := wallet.Subtract(123100.01210121)
		So(state, ShouldEqual, true)
		So(err, ShouldBeNil)
		So(wallet.Volume, ShouldEqual, 0.0)
	})
}

func panics() {
	panic("Goofy Gophers!")
}
