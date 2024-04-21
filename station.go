package main

import (
	"time"
)

type Station struct {
	Type         string
	Queue        chan *Car
	ServeTimeMin time.Duration
	ServeTimeMax time.Duration
}

func NewStation(stationType string, serveTimeMin, serveTimeMax time.Duration) *Station {
	return &Station{
		Type:         stationType,
		Queue:        make(chan *Car),
		ServeTimeMin: serveTimeMin,
		ServeTimeMax: serveTimeMax,
	}
}
