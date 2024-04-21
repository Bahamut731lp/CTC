package main

import (
	"time"
)

type CashRegister struct {
	Queue         chan *Car
	HandleTimeMin time.Duration
	HandleTimeMax time.Duration
}

func NewCashRegister(handleTimeMin, handleTimeMax time.Duration) *CashRegister {
	return &CashRegister{
		Queue:         make(chan *Car),
		HandleTimeMin: handleTimeMin,
		HandleTimeMax: handleTimeMax,
	}
}
