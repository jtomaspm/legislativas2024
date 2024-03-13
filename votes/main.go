package main

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"
	"votes/model"
	"votes/service"
)

func main() {
	//INPUT
	file, err := os.Open("../locations.csv")
	if err != nil {
		log.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	//OUTPUT
	fileOut, err := os.Create("../votes.csv")
	if err != nil {
		log.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer fileOut.Close()
	writer := csv.NewWriter(fileOut)
	defer writer.Flush()

	//PIPELINE
	locationsChannel := model.LoadLocations(scanner)
	territoryChannel := model.GetTerritoryResultsById(locationsChannel)
	service.SaveCsv(territoryChannel, writer)
}
