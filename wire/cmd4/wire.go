package main

import "github.com/google/wire"

var setOpts = wire.NewSet(
	"localhost:5000",
	"Default",
	"razor",
	"razor",
	"default",
	"service",
	wire.Struct(new(keystoneAdminAuthOptions), "Url", "UserDomainName", "Username", "Password", "ProjectDomainName", "ProjectName"))

func injectOptions() *keystoneAdminAuthOptions {
	wire.Build(setOpts)
	return &keystoneAdminAuthOptions{}
}
