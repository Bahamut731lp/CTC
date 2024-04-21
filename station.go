package main

import (
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
	MaxQueueTime time.Duration
}

func getNewStation(StationType string, ServeTimeMin, ServeTimeMax time.Duration) *Station {
	station := Station{}
	station.StationType = StationType
	station.ServeTimeMin = ServeTimeMin
	station.ServeTimeMax = ServeTimeMax
	station.MaxQueueTime = 0

	return &station
}
