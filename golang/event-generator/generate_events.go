package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	imo                 int
	timestamp           string
	lat                 float64
	lon                 float64
	berth_id            int
	navigational_status int
}

type Event struct {
	imo        int
	timestamp  string
	event_type string
}

func readPosition(line string) Position {
	row := strings.Split(line, ",")
	imo, err := strconv.Atoi(row[0])
	if err != nil {
		imo = 0
	}
	lat, err := strconv.ParseFloat(row[2], 64)
	if err != nil {
		lat = 0
	}
	lon, err := strconv.ParseFloat(row[3], 64)
	if err != nil {
		lon = 0
	}
	berth_id, err := strconv.ParseFloat(row[4], 32)
	if err != nil {
		berth_id = 0
	}
	navigational_status, _ := strconv.Atoi(row[5])
	if err != nil {
		navigational_status = 0
	}
	return Position{
		imo:                 imo,
		timestamp:           row[1],
		lat:                 lat,
		lon:                 lon,
		berth_id:            int(berth_id),
		navigational_status: navigational_status,
	}
}

func main() {
	file, err := os.Open("positions/2022-01_positions.csv")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	scanner.Scan() // skip header
	scanner.Scan()
	lastPosition := readPosition(scanner.Text())
	events := []Event{}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		position := readPosition(line)

		if position.berth_id != lastPosition.berth_id {
			events = append(events, Event{
				imo:        position.imo,
				timestamp:  position.timestamp,
				event_type: "berth_change",
			})
		}
		if position.navigational_status != lastPosition.navigational_status {
			events = append(events, Event{
				imo:        position.imo,
				timestamp:  position.timestamp,
				event_type: "status_change",
			})
		}
		lastPosition = position
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// write events to file
	file, err = os.Create("events/2022-01_events.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.WriteString("imo,timestamp,event_type\n")
	for _, event := range events {
		file.WriteString(fmt.Sprintf("%d,%s,%s\n", event.imo, event.timestamp, event.event_type))
	}
}
