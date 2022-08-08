// Package di is a wrapper of fx, which allows to use DI across the whole lifecycle.
package di

import (
	"go.uber.org/fx"
)

var headFxOptions = make([]fx.Option, 0, 1)
var globalFxOptions = make([]fx.Option, 0, 1)

// Provide a constructor method to create object for injection purpose.
//
// For example, we are going to inject `Some`:
//
//   type Some struct {
//     // -- snip --
//   }
//
//   func newSome() *Some {
//     // -- snip --
//   }
//
//   func init() {
//       di.Provide(newSome)
//   }
//
func Provide(newFn interface{}) {
	globalFxOptions = append(globalFxOptions, fx.Provide(newFn))
}

// Invoke a function with injected objects as parameters.
//
// For example, we want use `Some` object:
//
//     func init() {
//        di.Invoke(func(s *Some) {
//            // -- snip --
//        })
//     }
func Invoke(fn interface{}) {
	globalFxOptions = append(globalFxOptions, fx.Invoke(fn))
}

// Extract an injected object to a variable.
//
// For example, we can extrat injected `Some` object to a global variable:
//
//     type globals struct {
//     	some *Some
//     }
//
//     var g = &globals{};
//
//     func init() {
//     	di.Extract(&g);
//     }
func Extract(target interface{}) {
	globalFxOptions = append(globalFxOptions, fx.Extract(target))
}

// Populate allows to extract multiple injected objects to variables.
func Populate(targets ...interface{}) {
	globalFxOptions = append(globalFxOptions, fx.Populate(targets...))
}

// InvokeBefore invoke the given function before others.
func InvokeBefore(fn interface{}) {
	headFxOptions = append(headFxOptions, fx.Invoke(fn))
}

// Options of current project.
func Options() []fx.Option {
	results := make([]fx.Option, 0, len(headFxOptions)+len(globalFxOptions))
	results = append(results, headFxOptions...)
	results = append(results, globalFxOptions...)
	return results
}
