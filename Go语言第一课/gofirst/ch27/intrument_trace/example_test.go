package trace_test

import (
	. "github.com/andyron/gofirst/instrument_trace"
)

func a() {
	defer Trace()()
	b()
}

func b() {
	defer Trace()()
	c()
}

func c() {
	defer Trace()()
	d()
}

func d() {
	defer Trace()()
}

func ExampleTrace() {
	a()
	// Output:
	// g[00001]:    ->github.com/andyron/gofirst/instrument_trace_test.a
	// g[00001]:        ->github.com/andyron/gofirst/instrument_trace_test.b
	// g[00001]:            ->github.com/andyron/gofirst/instrument_trace_test.c
	// g[00001]:                ->github.com/andyron/gofirst/instrument_trace_test.d
	// g[00001]:                <-github.com/andyron/gofirst/instrument_trace_test.d
	// g[00001]:            <-github.com/andyron/gofirst/instrument_trace_test.c
	// g[00001]:        <-github.com/andyron/gofirst/instrument_trace_test.b
	// g[00001]:    <-github.com/andyron/gofirst/instrument_trace_test.a
}
