package main

import (
	"time"
)

type CarConfig struct {
	Count          int           `yaml:"count"`
	ArrivalTimeMin time.Duration `yaml:"arrival_time_min"`
	ArrivalTimeMax time.Duration `yaml:"arrival_time_max"`
}

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
