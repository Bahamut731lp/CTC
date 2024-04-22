package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/schollz/progressbar/v3"
)

type Config struct {
	Cars      CarConfig                `yaml:"cars"`
	Stations  map[string]StationConfig `yaml:"stations"`
	Registers Register                 `yaml:"registers"` //TODO: Kasy řešit něco jako semaforama i guess?
}

func getTimeInRange(min, max time.Duration) time.Duration {
	if min > max {
		min, max = max, min
	}

	rangeDuration := max - min
	randomDuration := time.Duration(rand.Int63n(int64(rangeDuration + 1)))
	finalDuration := min + randomDuration

	return finalDuration
}

func setCarsRoutine(config *Config, wg, wg2 *sync.WaitGroup, output chan *Car) {
	for range config.Cars.Count {
		wg.Add(1)
		wg2.Add(1)

		car := NewCar(time.Now())
		car.ArrivalTime = time.Now()
		time.Sleep(getTimeInRange(config.Cars.ArrivalTimeMin, config.Cars.ArrivalTimeMax))
		output <- car
	}

}

func main() {
	var wg, wg2 sync.WaitGroup
	config := getConfigFromFile("config.yml")

	noOfStations := 0
	noOfRegisters := config.Registers.Count

	log.Printf("Simulating gas station with %d cars.\n", config.Cars.Count)

	// Count the number of stations to allocate buffered channel
	for _, stationConfig := range config.Stations {
		noOfStations += stationConfig.Count
	}

	stations := []*Station{}
	registers := []*Register{}
	bar := progressbar.Default(int64(config.Cars.Count))

	inbound := make(chan *Car, config.Cars.Count)
	outbound := make(chan *Car, config.Cars.Count)

	for range noOfRegisters {
		register := Register{}
		register.HandleTimeMin = config.Registers.HandleTimeMin
		register.HandleTimeMax = config.Registers.HandleTimeMax
		register.TotalCars = 0
		register.TotalTime = 0
		register.MaxQueueTime = 0

		registers = append(registers, &register)
	}

	go setCarsRoutine(config, &wg, &wg2, inbound)

	// Fill up stations and registers
	for _type, config := range config.Stations {
		for i := range config.Count {
			station := Station{}
			station.StationType = fmt.Sprintf("%s %d", _type, i)
			station.ServeTimeMin = config.ServeTimeMin
			station.ServeTimeMax = config.ServeTimeMax
			station.MaxQueueTime = 0
			station.TotalCars = 0
			station.TotalTime = 0
			station.TotalQueueTime = 0

			stations = append(stations, &station)
		}
	}

	for _, station := range stations {
		go setStationRoutine(station, inbound, outbound, &wg)
	}

	for _, register := range registers {
		go setRegisterRoutine(register, outbound, &wg2, bar)
	}

	wg.Wait()
	wg2.Wait()

	log.Println("Simulation ended.")
	log.Println("Calculating statistics.")

	renderStationsStats(stations)
	renderRegistersStats(registers)
}
