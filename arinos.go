package arinos

type Arinos struct {
	LocalHost bool
}

type Options struct {
	Port int
}

type Option func(*Options)

func New(isLocalhost bool) (arinos *Arinos) {
	return &Arinos{
		LocalHost: isLocalhost,
	}
}

func Port(portNumber int) Option {
	return func(args *Options) {
		args.Port = portNumber
	}
}
