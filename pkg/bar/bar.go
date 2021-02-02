package bar

import "fmt"

// Foo ...
type Foo struct {
	verbosity int
	debug     bool
}

// Option ...
type Option func(*Foo) Option

// Options ...
func (f *Foo) Options(opts ...Option) (prevs []Option) {
	for _, opt := range opts {
		prevs = append(prevs, opt(f))
	}
	return prevs
}

// Verbosity ...
func Verbosity(v int) Option {
	return func(f *Foo) Option {
		var prev int
		prev, f.verbosity = f.verbosity, v
		fmt.Printf("verbosity now: %#v prev: %#v\n", f.verbosity, prev)
		return Verbosity(prev)
	}
}

// Debug ...
func Debug(d bool) Option {
	return func(f *Foo) Option {
		var prev bool
		prev, f.debug = f.debug, d
		fmt.Printf("debug now: %#v prev: %#v\n", f.debug, prev)
		return Debug(prev)
	}
}
