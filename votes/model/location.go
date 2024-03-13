package model

import (
	"bufio"
	"strings"
)

type Location struct {
	Id   string
	Type string
	Name string
}

func locationFromLine(line string) *Location {
	arr := strings.Split(line, ";")
	if len(arr) != 3 {
		return nil
	}
	loc := Location{
		Id:   arr[0],
		Type: arr[1],
		Name: arr[2],
	}
	return &loc
}

func LoadLocations(scanner *bufio.Scanner) <-chan Location {
	out := make(chan Location)
	go func() {
		for scanner.Scan() {
			line := scanner.Text()
			loc := locationFromLine(line)
			if loc != nil {
				out <- *loc
			}
		}
		close(out)
	}()
	return out
}
