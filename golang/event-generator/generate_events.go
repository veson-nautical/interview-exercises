package main

import (
	"fmt"
)

type VesselState struct {
	imo       int
	timestamp string
	port_id   int
	speed     int
}

type Event struct {
	imo        int
	timestamp  string
	event_type string
}

var vessel_states = []VesselState{
	{
		// IMO is a unique vessel identifier
		imo:       1,
		timestamp: "2020-01-01T00:00:00Z",
		// 0 means not in port
		port_id: 0,
		speed:   10,
	},
	{
		imo:       1,
		timestamp: "2020-01-02T00:00:00Z",
		port_id:   0,
		speed:     8,
	},
	{
		imo:       1,
		timestamp: "2020-01-03T00:00:00Z",
		port_id:   2,
		speed:     0,
	},
	{
		imo:       1,
		timestamp: "2020-01-03T00:00:00Z",
		port_id:   0,
		speed:     8,
	},
}

func main() {
	events := []Event{}
	previousState := vessel_states[0]
	for _, state := range vessel_states {
		if state.port_id != previousState.port_id {
			events = append(events, Event{
				imo:        state.imo,
				timestamp:  state.timestamp,
				event_type: "port_change",
			})
		}
		if state.speed != previousState.speed {
			events = append(events, Event{
				imo:        state.imo,
				timestamp:  state.timestamp,
				event_type: "speed_change",
			})
		}
		previousState = state
	}

	// write events
	for _, event := range events {
		fmt.Printf(fmt.Sprintf("%d,%s,%s\n", event.imo, event.timestamp, event.event_type))
	}
}
