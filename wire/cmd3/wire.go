// +build wireinject

package main

import "github.com/google/wire"

var Set = wire.NewSet(
	ProvideFoo,
	ProvideBar,
	wire.Struct(new(FooBar), "MyFoo", "MyBar"))

func injectFooBar() FooBar {
	wire.Build(Set)
	return FooBar{}
}
