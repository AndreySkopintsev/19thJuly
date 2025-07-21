module api

go 1.24.4

require (
	common v0.0.0-00010101000000-000000000000
	github.com/google/uuid v1.6.0
	github.com/gorilla/mux v1.8.1
)

require github.com/lpernett/godotenv v0.0.0-20230527005122-0de1d4c5ef5e // indirect

replace common => ../common
