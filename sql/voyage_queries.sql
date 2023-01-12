CREATE TABLE vessels (
    vessel_id int,
    imo int, -- ship's international maritime organization number
    vessel_name text,
    deadweight float,
)

CREATE TABLE ports (
    port_id int,
    port_name text,
)

CREATE TABLE voyages (
    voyage_id int,
    vessel_id int,
    commence_date timestamp
)

CREATE TABLE voyage_itineraries (
    voyage_id int,
    port_id int,
    seq_no int, -- the seqence number determines the order of the itinerary entries
    arrival_date timestamp,
    departure_date timestamp,
)

-- # Queries

-- Voyages commenced in the last year
-- Voyages for the vessel named ACACIA
-- How many arrivals will there be for each port in the next 2 weeks. (port_name, arrival_count)
-- The top 10 voyages with the most itinerary entries
-- Voyages that end at the port named HOUSTON
-- Voyages and their durations (arrival date of last itinerary - departure date of first itinerary)

-- # Extend the Schema

-- Store a vessel's length
-- Represent countries (country code & country name), ports belong to a country
-- Which columns might be worth indexing in the above tables?
-- What if vessel's could get reassigned in the middle of a voyage, how could you store that information?
    -- does your solution keep the history of the voyage?
