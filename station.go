package main

import (
	"math/rand"
	"time"
)

type Station struct {
	StationType  string
	ServeTimeMin time.Duration
	ServeTimeMax time.Duration
	IsAvailable  bool
	TotalCars    int
	TotalTime    time.Duration
}

func NewStation(StationType string, ServeTimeMin, ServeTimeMax time.Duration) *Station {
	station := Station{}
	station.StationType = StationType
	station.ServeTimeMin = ServeTimeMin
	station.ServeTimeMax = ServeTimeMax

	return &station
}

func (s *Station) Serve(car *Car) {
	serveTime := time.Duration(rand.Intn(int(s.ServeTimeMax-s.ServeTimeMin))) + s.ServeTimeMin
	car.StationTime = serveTime
	s.IsAvailable = false
	s.TotalCars++
	s.TotalTime += serveTime
	time.Sleep(serveTime)
	s.IsAvailable = true
}
