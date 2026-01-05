package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("processTruck", func(t *testing.T) {

		t.Run("Should load and unload a truck cargo", func(t *testing.T) {
			nt := &NormalTruck{id: "1"}
			et := &ElectricTruck{id: "2"}

			if err := processTruck(nt); err != nil {
				t.Fatalf("Error processing truck: %s", err)
			}

			if err := processTruck(et); err != nil {
				t.Fatalf("Error processing truck: %s", err)
			}

			// asserting
			if got := nt.cargo; got != 0 {
				t.Fatalf("Cargo should be 0, got %d", got)
			}

			if got := et.battery; got != -2 {
				t.Fatalf("Battery should be -2, got %d", got)
			}
		})
	})
}
