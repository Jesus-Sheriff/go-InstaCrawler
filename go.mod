module github.com/heroku/go-getting-started

// al indicar version 1.13 se usa la última minor de esa version mayor
// Versiones de Go soportadas: https://github.com/heroku/heroku-buildpack-go/blob/master/data.json#L3
go 1.13

require (
	github.com/struCoder/pmgo
	github.com/gorilla/mux
	github.com/joho/godotenv
	github.com/ahmdrz/goinsta
)
