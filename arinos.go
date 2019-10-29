package arinos

import (
	"fmt"
	"net/http"
	"net/http/fcgi"
	"time"
)

type Arinos struct {
	LocalHost bool
	Options   *Options
}

type Options struct {
	Port int
}

type Option func(*Options)

func New(isLocalhost bool, options ...Option) (arinos *Arinos) {
	args := &Options{
		Port: 8000,
	}
	for _, option := range options {
		option(args)
	}
	return &Arinos{
		LocalHost: isLocalhost,
		Options:   args,
	}
}

func Port(portNumber int) Option {
	return func(args *Options) {
		args.Port = portNumber
	}
}

// Mock
func (arinos *Arinos) StartServe() error {
	if arinos.LocalHost {
		portString := fmt.Sprintf(":%d", arinos.Options.Port)
		server := &http.Server{
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
			Addr:         portString,
		}
		return server.ListenAndServe()
	}
	return fcgi.Serve(nil, nil)
}
