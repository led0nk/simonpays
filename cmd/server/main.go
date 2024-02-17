package main

import (
	"errors"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

func main() {
	logger := zerolog.New(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339},
	).Level(zerolog.TraceLevel).With().Timestamp().Caller().Logger()

	logger.Info().Msg("Test")
	s := newServer(logger)
	router := mux.NewRouter()
	router.HandleFunc("/", s.homeHandler(logger)).Methods(http.MethodGet)

	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}
	logger.Info().Str("Key", "value").Msg("")
	s.userError("testError")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func (s *Server) homeHandler(logger zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("../../internal/templates/home.html", "../../internal/templates/header.html", "../../internal/templates/main.html")
		if err != nil {
			logger.Info().Msg("Test")
			logger.Error().Str("Error", "error")
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			logger.Error().Str("Error", "error")
			logger.Info().Msg("Test")
			return
		}
	}
}

type Server struct {
	logger zerolog.Logger
}

func newServer(logger zerolog.Logger) *Server {
	return &Server{
		logger: logger,
	}
}

func (s *Server) userError(errStr string) {
	err := errors.New("user")
	s.logger.Error().Err(err).Msg(errStr)
}
