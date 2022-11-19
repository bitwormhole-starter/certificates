package certificates

import (
	"embed"

	"github.com/bitwormhole/starter"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
)

const (
	theModuleName        = "bitwormhole.com/starter/certificates"
	theModuleVersion     = "v0.0.1"
	theModuleRevision    = 1
	theModuleResPath     = "src/main/resources"
	theModuleTestResPath = "src/test/resources"
)

////////////////////////////////////////////////////////////////////////////////

//go:embed "src/main/resources"
var theModuleResFS embed.FS

// InitModule ...
func InitModule(mb *application.ModuleBuilder) {

	mb.Name(theModuleName)
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)

	mb.Resources(collection.LoadEmbedResources(&theModuleResFS, theModuleResPath))

	// mb.OnMount(nil)

	mb.Dependency(starter.Module())
}

////////////////////////////////////////////////////////////////////////////////

//go:embed "src/test/resources"
var theModuleTestResFS embed.FS

// InitModuleTest ...
func InitModuleTest(mb *application.ModuleBuilder) {

	mb.Name(theModuleName + "#unit")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)

	mb.Resources(collection.LoadEmbedResources(&theModuleTestResFS, theModuleTestResPath))

	// mb.OnMount(nil)
	// mb.Dependency(starter.Module())
}

////////////////////////////////////////////////////////////////////////////////
