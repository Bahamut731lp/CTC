package main

import (
	"math/rand"
	"time"
)

type StationConfig struct {
	Count        int           `yaml:"count"`
	ServeTimeMin time.Duration `yaml:"serve_time_min"`
	ServeTimeMax time.Duration `yaml:"serve_time_max"`
}

type Station struct {
	StationType  string
	ServeTimeMin time.Duration `yaml:"serve_time_min"`
	ServeTimeMax time.Duration `yaml:"serve_time_max"`
	IsAvailable  bool
	TotalCars    int
	TotalTime    time.Duration
}

func getNewStation(StationType string, ServeTimeMin, ServeTimeMax time.Duration) *Station {
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
