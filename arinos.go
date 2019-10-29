package arinos

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
