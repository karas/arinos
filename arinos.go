package arinos

type Arinos struct {
	Options *Options
}

type Options struct {
	LocalHost bool
	Port      int
}

func New() (arinos *Arinos) {
	return &Arinos{}
}
