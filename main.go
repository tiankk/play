package main

import "github.com/tiankk/play/pkg/bar"

// DoSth ...
func DoSth(foo *bar.Foo, verbosity int) {
	prev := foo.Options(bar.Verbosity(verbosity), bar.Debug(true))
	defer foo.Options(prev...)

	// do something ...
}

func main() {
	DoSth(&bar.Foo{}, 3)
}
