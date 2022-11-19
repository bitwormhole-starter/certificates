package lib

import "bitwormhole.com/starter/certificates"

type chainImpl struct {
	parent certificates.Chain
	cert   certificates.Certificate
}

func (inst *chainImpl) _Impl() certificates.Chain {
	return inst
}

func (inst *chainImpl) GetParent() certificates.Chain {
	return inst.parent
}

func (inst *chainImpl) GetCertificate() certificates.Certificate {
	return inst.cert
}
