# Events generator

This is a simple go script which converts lat/lon vessel coordinate date to vessel events
by iterating line by line over the positions file and detecting changes in the berth or status columns 
and writing out the corresponding berth_changed or status_changed event.

## To Run

`go run generate_events.go`

## Tasks

- run over all files in positions dir
- use a csv library instead of manual parsing
- general code clean ups / bigfixes
- process input files in parallel
- handle berth entry / exit instead of just change
- make the code more generic to handle arbitrary changing columns
- pass the remaining positions columns to the events table
- (advanced) brainstorm how to handle boundaries between positions files

