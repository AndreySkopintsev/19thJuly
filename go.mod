module main

go 1.24.4

replace api => ./api

require (
	api v0.0.0-00010101000000-000000000000
	common v0.0.0-00010101000000-000000000000
)

require (
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorilla/mux v1.8.1 // indirect
	github.com/lpernett/godotenv v0.0.0-20230527005122-0de1d4c5ef5e // indirect
)

replace common => ./common
