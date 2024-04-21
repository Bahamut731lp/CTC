package main

import (
	"time"
)

type Car struct {
	ArrivalTime  time.Time
	QueueTime    time.Duration
	StationTime  time.Duration
	RegisterTime time.Duration
}

func NewCar(ArrivalTime time.Time, QueueTime, StationTime, RegisterTime time.Duration) *Car {
	car := Car{}
	car.ArrivalTime = ArrivalTime
	car.QueueTime = QueueTime
	car.StationTime = StationTime
	car.RegisterTime = RegisterTime

	return &car
}
