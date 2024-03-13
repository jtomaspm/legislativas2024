package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("dump.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	fileOut, err := os.OpenFile("../locations.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer fileOut.Close()

	i := 0
	content := ""
	for scanner.Scan() {
		i = i + 1
		line := scanner.Text()
		if i == 1 {
			content += line
		} else if i == 2 {
			content += ";"
		} else if i == 3 {
			content += line[1:len(line)-1] + "\n"
			i = 0
			_, err = fileOut.WriteString(content)
			if err != nil {
				fmt.Println("Error writing to file:", err)
				os.Exit(1)
			}
			content = ""
		} else {
			fmt.Println("Error, skill issue")
			os.Exit(1)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}
}
