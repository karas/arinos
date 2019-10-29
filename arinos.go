package arinos

type Arinos struct {
	LocalHost bool
}

type Options struct {
	Port int
}

func New(isLocalhost bool) (arinos *Arinos) {
	return &Arinos{
		LocalHost: isLocalhost,
	}
}
