// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package cfglibcerts

import (
	certificates0xb7c330 "bitwormhole.com/starter/certificates"
	lib0x47d8cf "bitwormhole.com/starter/certificates/lib"
	roots0x5e4213 "bitwormhole.com/starter/certificates/lib/roots"
	x509solution0x7650e0 "bitwormhole.com/starter/certificates/lib/solutions/x509solution"
	application0x67f6c5 "github.com/bitwormhole/starter/application"
	markup0x23084a "github.com/bitwormhole/starter/markup"
)

type pComResRoots struct {
	instance *roots0x5e4213.ResRoots
	 markup0x23084a.Component `class:"certificate-root-registry"`
	Context application0x67f6c5.Context `inject:"context"`
	Enabled bool `inject:"${certificates.roots-enabled}"`
	NameList string `inject:"${certificates.roots}"`
}


type pComRootCertManager struct {
	instance *lib0x47d8cf.RootCertManager
	 markup0x23084a.Component `id:"certificate-root-manager"`
	Solutions certificates0xb7c330.SolutionManager `inject:"#certificate-solution-manager"`
	Regs []certificates0xb7c330.RootRegistry `inject:".certificate-root-registry"`
}


type pComRsaX509Solution struct {
	instance *x509solution0x7650e0.RsaX509Solution
	 markup0x23084a.Component `class:"certificate-solution-registry"`
	Roots certificates0xb7c330.RootManager `inject:"#certificate-root-manager"`
}


type pComCertSolutionManager struct {
	instance *lib0x47d8cf.CertSolutionManager
	 markup0x23084a.Component `id:"certificate-solution-manager"`
	Roots certificates0xb7c330.RootManager `inject:"#certificate-root-manager"`
	Regs []certificates0xb7c330.SolutionRegistry `inject:".certificate-solution-registry"`
}

