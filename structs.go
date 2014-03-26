package main

import (
	"errors"
	"sync"
)

type User struct {
	Name string
	Id   int
}

type Currency struct {
	Name string
}

type ExchangeType struct {
	Direction uint
}

type Exchange struct {
	Type           *ExchangeType
	Rate           float64
	TargetCurrency *Currency
	OriginCurrency *Currency
}

type Transaction struct {
	Volume float64
	Origin *Wallet
	Target *Wallet
	Done   bool
}

type Wallet struct {
	Owner    *User
	Volume   float64
	Currency *Currency
	Mutex    *sync.Mutex
}

func NewUser(name string, id int) *User {
	return &User{
		Name: name,
		Id:   id,
	}
}

func NewCurrency(name string) *Currency {
	return &Currency{
		Name: name,
	}
}

func NewTransaction(volume float64, origin *Wallet, target *Wallet) *Transaction {
	return &Transaction{
		Volume: volume,
		Origin: origin,
		Target: target,
		Done:   false,
	}
}

func NewWallet(owner *User, volume float64, currency *Currency) *Wallet {
	return &Wallet{
		Owner:    owner,
		Currency: currency,
		Volume:   volume,
		Mutex:    &sync.Mutex{},
	}
}

func NewExchange(exchangeType *ExchangeType, rate float64, targetCurrency *Currency, originCurrency *Currency) *Exchange {
	return &Exchange{
		Type:           exchangeType,
		Rate:           rate,
		TargetCurrency: targetCurrency,
		OriginCurrency: originCurrency,
	}
}

func NewExchangeType(direction uint) *ExchangeType {
	return &ExchangeType{
		Direction: direction,
	}
}

func (w *Wallet) Subtract(volume float64) (bool, error) {
	w.Mutex.Lock()
	defer w.Mutex.Unlock()

	delta := w.Volume - volume
	if delta >= 0 {
		w.Volume = delta
		return true, nil
	} else {
		return false, errors.New("Negative delta on subtraction")
	}
}

func (w *Wallet) Add(volume float64) (bool, error) {
	w.Mutex.Lock()
	defer w.Mutex.Unlock()

	delta := w.Volume + volume
	if delta >= 0 {
		w.Volume = delta
		return true, nil
	} else {
		return false, errors.New("Negative delta on addition")
	}
}

func (t *Transaction) Rollback() error {
	_, err := t.Origin.Add(t.Volume)
	return err
}

func (t *Transaction) Commit() (bool, error) {
	if *t.Origin.Currency != *t.Target.Currency {
		return false, WrongCurrencyError{
			t.Origin.Currency,
			t.Target.Currency,
		}
	}
	state, err := t.Origin.Subtract(t.Volume)
	if err != nil {
		return state, err
	}
	state, err = t.Target.Add(t.Volume)
	if err != nil {
		return state, t.Rollback()
	} else {
		t.Done = true
	}
	return state, err
}
