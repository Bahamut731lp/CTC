package main

import (
	"log"
	"sync"
)

type Config struct {
	Cars      CarConfig                `yaml:"cars"`
	Stations  map[string]StationConfig `yaml:"stations"`
	Registers Register                 `yaml:"count"` //TODO: Kasy řešit něco jako semaforama i guess?
}

func main() {
	config := getConfigFromFile("config.yml")

	// Tohle funguje v postatě jako semafor
	// registers := make(chan int, config.Registers.Count)

	stations := sync.Pool{}
	//queue := make([]Car, 0)

	//	TODO: Create goroutine to add new cars to queue and one to check for available stations in pool
	log.Printf("Simulating gas station with %d cars.\n", config.Cars.Count)
	// Vytváření nových stanic
	for stationType, stationConfig := range config.Stations {
		log.Printf("Adding %d stand(s) with fuel type \"%s\"\n", stationConfig.Count, stationType)
		for range stationConfig.Count {
			station := getNewStation(stationType, stationConfig.ServeTimeMin, stationConfig.ServeTimeMax)
			stations.Put(station)
		}
	}
}
