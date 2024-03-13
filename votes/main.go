package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"votes/model"
)

func getUrl(territoryKey string) string {
	return "https://www.legislativas2024.mai.gov.pt/frontend/data/TerritoryResults?territoryKey=" + territoryKey + "&electionId=AR"
}

func main() {
	url := getUrl("LOCAL-030000")
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var data model.TerritoryResults
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	fmt.Println("Parsed data:", data.CurrentResults.BlankVotes)
}
