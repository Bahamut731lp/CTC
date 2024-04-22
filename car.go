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
	ArrivalTime       time.Time
	StationQueueTime  time.Duration
	StationTime       time.Duration
	RegisterTime      time.Duration
	RegisterQueueTime time.Duration
}

func NewCar(ArrivalTime time.Time) *Car {
	car := Car{}
	car.ArrivalTime = ArrivalTime
	car.StationQueueTime = 0
	car.RegisterQueueTime = 0
	car.StationTime = 0
	car.RegisterTime = 0

	return &car
}
