# Events generator

This is a simple go script which converts vessel status updates to vessel events (timed changes in state)
by iterating line by line over the states data and detecting changes in the port or speed columns 
and writing out the corresponding port_changed or speed_changed event.

## To Run

`go run generate_events.go`

## Specific Tasks

1. The code only detects port_change events, update it to instead detect port_entry and port_exit events.
2. The code assumes all input events are for a single vessel (IMO). Update it to handle multiple vessel events at once.
    - Can you clean the code up now to split it into several functions to improve readability
3. (Conceptual) How would you handle reading/writing the input and output to files
    - What format(s) might you use?
    - How would you implement the reading/writing (what libraries might you use)
    - If there was a lot of data you might need to read/write to and from multiple files instead of one large file. Assuming a new batch of state files came in for all vessels every day. How would you handle that?
4. (Advanced) can you update the implementation to take advantage of multithreading.

## General Tasks

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
