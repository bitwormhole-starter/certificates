package x509solution

import (
	"fmt"

	"bitwormhole.com/starter/certificates"
	"github.com/bitwormhole/starter/markup"
)

// RsaX509Solution 。。。
type RsaX509Solution struct {
	markup.Component `class:"certificate-solution-registry"`

	Roots certificates.RootManager `inject:"#certificate-root-manager"`
}

func (inst *RsaX509Solution) _Impl() (certificates.Solution, certificates.SolutionRegistry) {
	return inst, inst
}

// GetRegistration ...
func (inst *RsaX509Solution) GetRegistration() *certificates.SolutionRegistration {
	info := inst.GetInfo()
	return &certificates.SolutionRegistration{
		Info:     *info,
		Solution: inst,
	}
}

// Load ...
func (inst *RsaX509Solution) Load(raw *certificates.Raw) ([]certificates.Certificate, error) {
	if raw == nil {
		return nil, fmt.Errorf("raw data is nil")
	}
	loader := myRsaX509Loader{solution: inst}
	return loader.load(raw.Data)
}

// Verify ...
func (inst *RsaX509Solution) Verify(chain certificates.Chain, op *certificates.VerifyOptions) error {
	veri := x509veri{}
	veri.chain = chain
	veri.roots = inst.Roots
	if op != nil {
		veri.options = *op
	}
	return veri.verify()
}

// GetInfo ...
func (inst *RsaX509Solution) GetInfo() *certificates.SolutionInfo {
	return &certificates.SolutionInfo{
		Algorithm:   "rsa",
		Format:      "x509",
		ContentType: "pem",
	}
}
