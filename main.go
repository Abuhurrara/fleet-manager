package main

import (
	"errors"
	"fmt"
	"log"
)

var (
	ErrNotImplemented = errors.New("not implemented")
	ErrNotFound       = errors.New("not found")
)

type Truck interface {
	LoadCargo() error
	UnloadCargo() error
}

type NormalTruck struct {
	id    string
	cargo int
}

func (t NormalTruck) LoadCargo() error {
	t.cargo += 1
	return nil
}

func (t NormalTruck) UnloadCargo() error {
	t.cargo = 0
	return nil
}

type ElectricTruck struct {
	id      string
	cargo   int
	battery int
}

func (e ElectricTruck) LoadCargo() error {
	e.cargo += 1
	e.battery -= 1
	return nil
}

func (e ElectricTruck) UnloadCargo() error {
	e.cargo = 0
	e.battery += -1
	return nil
}

func processTruck(truck Truck) error {
	fmt.Printf("Processing truck %+v \n", truck)
	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("error loading cargo: %w", err)
	}

	if err := truck.UnloadCargo(); err != nil {
		return fmt.Errorf("error unloading cargo: %w", err)
	}

	return nil
}

func main() {

	if err := processTruck(NormalTruck{id: "1"}); err != nil {
		log.Fatalf("Error processing truck: %s", err)
	}

	if err := processTruck(ElectricTruck{id: "2"}); err != nil {
		log.Fatalf("Error processing truck: %s", err)
	}
}
