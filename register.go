package main

import "time"

type Register struct {
	Count         int           `yaml:"count"`
	HandleTimeMin time.Duration `yaml:"handle_time_min"`
	HandleTimeMax time.Duration `yaml:"handle_time_max"`
}
