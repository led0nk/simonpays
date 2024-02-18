package main

import (
	"os"
	"time"

	"github.com/gorilla/mux"
	v1 "github.com/led0nk/webshop/api/v1"
	"github.com/rs/zerolog"
)

func main() {
	logger := zerolog.New(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339},
	).Level(zerolog.TraceLevel).With().Timestamp().Caller().Logger()
	// create opts and server
	opts := v1.NewOpts("localhost:8080", mux.NewRouter())
	server := v1.NewServer(logger, opts)

	// run server
	server.ServeHTTP()
}
