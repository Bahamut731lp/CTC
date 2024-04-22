package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/olekukonko/tablewriter"
)

type Register struct {
	Count         int           `yaml:"count"`
	HandleTimeMin time.Duration `yaml:"handle_time_min"`
	HandleTimeMax time.Duration `yaml:"handle_time_max"`
	TotalCars     int
	TotalTime     time.Duration
	MaxQueueTime  time.Duration
}

func setRegisterRoutine(register *Register, input chan *Car, wg *sync.WaitGroup) {
	for car := range input {
		//queueTime := time.Since(car.ArrivalTime)
		serveTime := getTimeInRange(register.HandleTimeMin, register.HandleTimeMax)
		car.RegisterTime = serveTime
		car.RegisterQueueTime = time.Since(car.ArrivalTime.Add(car.StationQueueTime).Add(car.StationTime))

		register.TotalCars += 1
		register.TotalTime += serveTime

		if register.MaxQueueTime < car.RegisterQueueTime {
			register.MaxQueueTime = car.StationQueueTime
		}

		time.Sleep(serveTime)
		wg.Done()
	}
}

func renderRegistersStats(registers []*Register) {
	table := tablewriter.NewWriter(os.Stdout)
	// Define table headers
	table.SetHeader([]string{"Type", "Total Cars", "Total Time"})

	// Add table rows
	for i, register := range registers {
		table.Append([]string{fmt.Sprint(i), fmt.Sprint(register.TotalCars), fmt.Sprint(register.TotalTime)})
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
