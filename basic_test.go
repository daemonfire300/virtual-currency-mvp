package main

import (
	//"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"math"
	"sync"
	"testing"
)

func TestAssertions(t *testing.T) {
	user := NewUser("Frank", 1234)
	currency := NewCurrency("FruitCoins")
	currencyDupe := NewCurrency("FruitCoins")
	wallet := NewWallet(user, 123100.0, currency)
	walletC := NewWallet(user, 123100.0, currency)
	walletD := NewWallet(user, 123100.0, currencyDupe)
	walletE := NewWallet(user, 2000.0, currency)
	walletF := NewWallet(user, 1000.0, currencyDupe)

	currencyB := NewCurrency("FruitCoinsCS")
	walletB := NewWallet(user, 123100.0, currencyB)

	//chan1 := make(chan int)
	//chan2 := make(chan int)

	Convey("Relations tests", t, func() {
		So(nil, ShouldBeNil)
		So(wallet.Currency, ShouldEqual, currency)
		So(wallet.Owner, ShouldEqual, user)
		So(wallet.Volume, ShouldBeGreaterThan, 0.0)
		So(user.Name, ShouldEqual, "Frank")
	})

	Convey("Wallet tests", t, func() {
		So(*wallet.Currency, ShouldResemble, *currencyDupe)
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
		transaction := NewTransaction(100.0, walletD, walletC)
		state, err := transaction.Commit()
		So(state, ShouldEqual, true)
		So(err, ShouldBeNil)
		newValue := 123100.0 - 100.0
		So(walletD.Volume, ShouldEqual, newValue)
	})

	Convey("Transaction Right concurrent commits", t, func() {
		var wg sync.WaitGroup

		testTrans := func() {
			defer wg.Done()
			transactionK := NewTransaction(1.5, walletE, walletF)
			transactionK.Commit()
		}

		for i := 0; i < 1000; i++ {
			wg.Add(1)
			go testTrans()
		}
		wg.Wait()
		So(walletF.Volume+walletE.Volume, ShouldEqual, 3000.0)
		So(math.Abs(walletF.Volume-walletE.Volume), ShouldEqual, 2000.0)
	})
}
