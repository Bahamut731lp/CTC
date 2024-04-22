package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/olekukonko/tablewriter"
)

type Config struct {
	Cars      CarConfig                `yaml:"cars"`
	Stations  map[string]StationConfig `yaml:"stations"`
	Registers Register                 `yaml:"count"` //TODO: Kasy řešit něco jako semaforama i guess?
}

func process(car *Car, station Station, stations *sync.Pool) {
	// Serving at the station
	queueTime := time.Since(car.ArrivalTime)
	serveTime := getTimeInRange(station.ServeTimeMin, station.ServeTimeMax)
	station.TotalCars += 1
	station.TotalTime += serveTime

	car.QueueTime = queueTime
	car.StationTime = serveTime
	if car.QueueTime > station.MaxQueueTime {
		station.MaxQueueTime = car.QueueTime
	}

	log.Printf("Car is done\n")
	stations.Put(station)
}

func getTimeInRange(min, max time.Duration) time.Duration {
	return min + time.Duration(rand.Int63n(int64(max-min+1)))
}

func main() {
	var wg sync.WaitGroup
	config := getConfigFromFile("config.yml")

	//arrivalTimeMin := config.Cars.ArrivalTimeMin
	//arrivalTimeMax := config.Cars.ArrivalTimeMax
	noOfStations := 0
	noOfRegisters := config.Registers.Count

	log.Printf("Simulating gas station with %d cars.\n", config.Cars.Count)

	// Count the number of stations to allocate buffered channel
	for _, stationConfig := range config.Stations {
		noOfStations += stationConfig.Count
	}

	stations := []*Station{}
	registers := make(chan *Register, noOfRegisters)

	inbound := make(chan *Car, config.Cars.Count)
	//outbound := make(chan *Car, config.Cars.Count)

	for range noOfRegisters {
		register := Register{}
		register.HandleTimeMin = config.Registers.HandleTimeMin
		register.HandleTimeMax = config.Registers.HandleTimeMax

		registers <- &register
	}

	// Spawning cars
	for range config.Cars.Count {
		wg.Add(1)

		car := NewCar(time.Now())
		car.ArrivalTime = time.Now()
		inbound <- car
	}

	// Fill up stations and registers
	for _type, config := range config.Stations {
		station := Station{}
		station.StationType = _type
		station.ServeTimeMin = config.ServeTimeMin
		station.ServeTimeMax = config.ServeTimeMax

		stations = append(stations, &station)

		go func(station *Station, input chan *Car, wg *sync.WaitGroup) {
			for car := range input {
				queueTime := time.Since(car.ArrivalTime)
				serveTime := getTimeInRange(station.ServeTimeMin, station.ServeTimeMax)

				station.TotalCars += 1
				station.TotalTime += serveTime

				car.QueueTime = queueTime
				car.StationTime = serveTime
				if car.QueueTime > station.MaxQueueTime {
					station.MaxQueueTime = car.QueueTime
				}

				time.Sleep(serveTime)
				wg.Done()
			}
		}(&station, inbound, &wg)
	}

	wg.Wait()

	log.Println("Simulation ended.")
	log.Println("Calculating statistics.")

	table := tablewriter.NewWriter(os.Stdout)
	// Define table headers
	table.SetHeader([]string{"Type", "Total Cars", "Total Time", "Average Queue Time", "Maximum Queue Time"})

	// Add table rows
	for i := 0; i < len(stations); i++ {
		station := stations[i]
		table.Append([]string{station.StationType, fmt.Sprint(station.TotalCars), fmt.Sprint(station.TotalTime), fmt.Sprint(station.TotalTime / time.Duration(station.TotalCars)), fmt.Sprint(station.MaxQueueTime)})
	}

	// Set alignment for columns
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	// Set the table format
	table.SetBorder(false)
	table.SetRowLine(true)
	table.SetAutoWrapText(false)

	// Render the table
	table.Render()

}
