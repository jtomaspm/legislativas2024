package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"votes/model"
)

func main() {
	file, err := os.Open("../locations.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	locationsChannel := model.LoadLocations(scanner)

	territoryChannel := model.GetTerritoryResultsById(locationsChannel)

	for n := range territoryChannel {
		log.Println(n)
	}
}
