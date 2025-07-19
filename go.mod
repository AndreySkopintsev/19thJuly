module main

go 1.24.4

replace api => ./api

require api v0.0.0-00010101000000-000000000000

require github.com/gorilla/mux v1.8.1 // indirect
