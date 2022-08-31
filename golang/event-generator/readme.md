# Events generator

This is a simple go script which converts lat/lon vessel coordinate date to vessel events
by iterating line by line over the positions file and detecting changes in the berth or status columns 
and writing out the corresponding berth_changed or status_changed event.

## To Run

`go run generate_events.go`

## Tasks

- general improvements
    - add comments
    - don't re-process the first position
    - extract a test for the functionality
- improve event detection
    - pull it out into a separate function
    - create berth entry / exit events instead of just change events
    - copy other positions fields into the events struct
    - handle different column changes more generically
- I/O
    - write the events to a file instead of just logging
    - read positions from the csv files instead of the hard-coded example
    - Extra credit: read from files in parallel
    - (advanced) brainstorm how to handle boundaries between the end of one positions file and the start of the next
