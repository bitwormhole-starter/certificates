package main

import (
	"bitwormhole.com/starter/certificates"
	"bitwormhole.com/starter/certificates/gen/unit4libcerts"
	"bitwormhole.com/starter/certificates/libcerts"
	"github.com/bitwormhole/starter"
	"github.com/bitwormhole/starter/application"
)

// Module ...
func Module() application.Module {
	mb := &application.ModuleBuilder{}
	certificates.InitModuleTest(mb)
	mb.Dependency(libcerts.Module())
	mb.OnMount(unit4libcerts.ExportConfig)
	return mb.Create()
}

func main() {
	mod := Module()
	i := starter.InitApp()
	i.UseMain(mod)
	i.Run()
}
