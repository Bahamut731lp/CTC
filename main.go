package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/schollz/progressbar/v3"
)

type Config struct {
	Cars      CarConfig                `yaml:"cars"`
	Stations  map[string]StationConfig `yaml:"stations"`
	Registers Register                 `yaml:"count"` //TODO: Kasy řešit něco jako semaforama i guess?
}

func getTimeInRange(min, max time.Duration) time.Duration {
	return time.Duration(rand.Intn(int(max)-int(min+1)) + int(min))
}

func addNewCarsToQueue(config *Config, queue chan<- Car, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	max := config.Cars.Count
	arrivalTimeMin := config.Cars.ArrivalTimeMin
	arrivalTimeMax := config.Cars.ArrivalTimeMax

	log.Printf("Started adding %d cars into the simulation\n", max)
	bar := progressbar.Default(int64(max))

	for i := 0; i < max; i++ {
		timeToSleep := getTimeInRange(arrivalTimeMin, arrivalTimeMax)
		time.Sleep(timeToSleep)

		arrival := time.Now()
		car := NewCar(arrival)
		queue <- *car
		bar.Add(1)
	}
}

func getStationRunning(station *Station, queue <-chan Car, wg *sync.WaitGroup) {
	for car := range queue {
		wg.Add(1)
		serveTime := getTimeInRange(station.ServeTimeMin, station.ServeTimeMax)

		car.StationTime = serveTime
		car.QueueTime = time.Since(car.ArrivalTime)
		car.StationTime = serveTime

		station.TotalCars++
		station.TotalTime += serveTime

		if station.MaxQueueTime < car.QueueTime {
			station.MaxQueueTime = car.QueueTime
		}

		time.Sleep(serveTime)
		wg.Done()
	}
}

func main() {
	config := getConfigFromFile("config.yml")
	var wg sync.WaitGroup

	queue := make(chan Car)
	stations := []*Station{}

	log.Printf("Simulating gas station with %d cars.\n", config.Cars.Count)

	go addNewCarsToQueue(config, queue, &wg)

	// Vytváření nových stanic
	for stationType, stationConfig := range config.Stations {
		log.Printf("Adding %d %s stand(s)\n", stationConfig.Count, stationType)
		for range stationConfig.Count {
			station := getNewStation(stationType, stationConfig.ServeTimeMin, stationConfig.ServeTimeMax)
			stations = append(stations, station)

			go getStationRunning(station, queue, &wg)
		}
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

	println(fmt.Sprint(len(stations)), fmt.Sprint(len(queue)))
}
