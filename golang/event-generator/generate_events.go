package main

import (
	"fmt"
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

var positions = []Position{
	{
		imo:                 1,
		timestamp:           "2020-01-01T00:00:00Z",
		lat:                 29.0,
		lon:                 -140.0,
		berth_id:            0,
		navigational_status: 0,
	},
	{
		imo:                 1,
		timestamp:           "2020-01-02T00:00:00Z",
		lat:                 30.0,
		lon:                 -130.0,
		berth_id:            0,
		navigational_status: 3,
	},
	{
		imo:                 1,
		timestamp:           "2020-01-03T00:00:00Z",
		lat:                 40.0,
		lon:                 -140.0,
		berth_id:            2,
		navigational_status: 1,
	},
}

func main() {
	events := []Event{}
	previousPosition := positions[0]
	for _, position := range positions {
		if position.berth_id != previousPosition.berth_id {
			events = append(events, Event{
				imo:        position.imo,
				timestamp:  position.timestamp,
				event_type: "berth_change",
			})
		}
		if position.navigational_status != previousPosition.navigational_status {
			events = append(events, Event{
				imo:        position.imo,
				timestamp:  position.timestamp,
				event_type: "status_change",
			})
		}
		previousPosition = position
	}

	// write events
	for _, event := range events {
		fmt.Printf(fmt.Sprintf("%d,%s,%s\n", event.imo, event.timestamp, event.event_type))
	}
}
