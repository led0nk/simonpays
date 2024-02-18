package v1

import (
	"errors"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

type Options struct {
	Address string
	Handler http.Handler
}

type Server struct {
	logger zerolog.Logger
	opts   *Options
}

func NewOpts(address string, handler http.Handler) *Options {
	return &Options{
		Address: address,
		Handler: handler,
	}
}

func NewServer(logger zerolog.Logger, opts *Options) *Server {
	return &Server{
		logger: logger,
		opts:   opts,
	}
}

func (s *Server) ServeHTTP() {
	router := mux.NewRouter()

	// routes:
	router.HandleFunc("/", s.homeHandler()).Methods(http.MethodGet)
	err := errors.New("user cannot be found")
	s.logger.Error().Err(errors.New("user")).Msg(err.Error())
	s.logger.Error().Err(err).Msg("user")
	server := &http.Server{
		Addr:    s.opts.Address,
		Handler: router,
	}

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// handler for route "/"
func (s *Server) homeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("../../internal/templates/home.html", "../../internal/templates/header.html", "../../internal/templates/main.html")
		if err != nil {
			s.logger.Error().Err(errors.New("template")).Msg(err.Error())
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			s.logger.Error().Err(errors.New("template")).Msg(err.Error())
			return
		}
	}
}
