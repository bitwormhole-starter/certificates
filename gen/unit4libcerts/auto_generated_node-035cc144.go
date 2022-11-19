// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package unit4libcerts

import (
	certificates0xb7c330 "bitwormhole.com/starter/certificates"
	unit0x9449e7 "bitwormhole.com/starter/certificates/src/test/golang/unit"
	application0x67f6c5 "github.com/bitwormhole/starter/application"
	markup0x23084a "github.com/bitwormhole/starter/markup"
)

type pComTestPoint1 struct {
	instance *unit0x9449e7.TestPoint1
	 markup0x23084a.Component `class:"life"`
	Context application0x67f6c5.Context `inject:"context"`
	Solutions certificates0xb7c330.SolutionManager `inject:"#certificate-solution-manager"`
	CertFile string `inject:"${test.cert.file}"`
}

