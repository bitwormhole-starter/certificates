package libcerts

import (
	"bitwormhole.com/starter/certificates"
	"bitwormhole.com/starter/certificates/gen/cfglibcerts"
	"github.com/bitwormhole/starter/application"
)

// Module ... 导出模块
func Module() application.Module {
	mb := &application.ModuleBuilder{}
	certificates.InitModule(mb)
	mb.OnMount(cfglibcerts.ExportConfig)
	return mb.Create()
}
