package bom

import (
	"fmt"
	pprint "github.com/NubeIO/bom-api/print"
	"testing"
)

func TestClient_Observations(t *testing.T) {
	search, err := client.ObservationByTown("helensburgh")
	if err != nil {
		fmt.Print(err)
	}
	pprint.PrintJSON(search)
}

func TestClient_ObservationsByZip(t *testing.T) {
	search, err := client.ObservationByZip("2508")
	if err != nil {
		fmt.Print(err)
	}
	pprint.PrintJSON(search)
}
