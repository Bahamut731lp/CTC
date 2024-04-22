package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/olekukonko/tablewriter"
)

type StationConfig struct {
	Count        int           `yaml:"count"`
	ServeTimeMin time.Duration `yaml:"serve_time_min"`
	ServeTimeMax time.Duration `yaml:"serve_time_max"`
}

type Station struct {
	StationType    string
	ServeTimeMin   time.Duration `yaml:"serve_time_min"`
	ServeTimeMax   time.Duration `yaml:"serve_time_max"`
	IsAvailable    bool
	TotalCars      int
	TotalTime      time.Duration
	TotalQueueTime time.Duration
	MaxQueueTime   time.Duration
}

func setStationRoutine(station *Station, input chan *Car, output chan *Car, wg *sync.WaitGroup) {
	for car := range input {
		queueTime := time.Since(car.ArrivalTime)
		serveTime := getTimeInRange(station.ServeTimeMin, station.ServeTimeMax)

		station.TotalCars += 1
		station.TotalTime += serveTime
		station.TotalQueueTime += queueTime

		car.StationQueueTime = queueTime
		car.StationTime = serveTime

		if station.MaxQueueTime < car.StationQueueTime {
			station.MaxQueueTime = car.StationQueueTime
		}

		time.Sleep(serveTime)
		wg.Done()

		select {
		case output <- car:
			// Successfully moved the value from ch1 to ch2
		default:
			// ch2 is not ready to receive, skip this value for now
			fmt.Println("ch2 is not ready to receive, skipping value.")
		}
	}
}

func renderStationsStats(stations []*Station) {
	table := tablewriter.NewWriter(os.Stdout)
	// Define table headers
	table.SetHeader([]string{"Type", "Total Cars", "Total Time", "Average Queue Time", "Maximum Queue Time"})

	// Add table rows
	for i, station := range stations {
		table.Append([]string{fmt.Sprintf("%s %d", station.StationType, i), fmt.Sprint(station.TotalCars), fmt.Sprint(station.TotalTime), fmt.Sprint(station.TotalQueueTime / time.Duration(station.TotalCars)), fmt.Sprint(station.MaxQueueTime)})
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
