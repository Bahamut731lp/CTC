package main

import (
	"time"
)

type Car struct {
	ID          int
	ArrivalTime time.Time
	StartTime   time.Time
	EndTime     time.Time
}

func NewCar(id int, arrivalTime time.Time) *Car {
	return &Car{
		ID:          id,
		ArrivalTime: arrivalTime,
	}
}
