package lib

import (
	"fmt"

	"bitwormhole.com/starter/certificates"
	"github.com/bitwormhole/starter/markup"
)

// CertSolutionManager implements certificates.SolutionManager
type CertSolutionManager struct {
	markup.Component `id:"certificate-solution-manager"`

	Roots certificates.RootManager        `inject:"#certificate-root-manager"`
	Regs  []certificates.SolutionRegistry `inject:".certificate-solution-registry"`

	all []certificates.Solution
}

func (inst *CertSolutionManager) _Impl() certificates.SolutionManager {
	return inst
}

func (inst *CertSolutionManager) loadAll() []certificates.Solution {
	src := inst.Regs
	dst := make([]certificates.Solution, 0)
	for _, sr1 := range src {
		sr2 := sr1.GetRegistration()
		solution := sr2.Solution
		dst = append(dst, solution)
	}
	return dst
}

// ListAll ...
func (inst *CertSolutionManager) ListAll() []certificates.Solution {
	src := inst.all
	if src == nil {
		src = inst.loadAll()
		inst.all = src
	}
	dst := make([]certificates.Solution, 0)
	dst = append(dst, src...)
	return dst
}

// Load ...
func (inst *CertSolutionManager) Load(raw *certificates.Raw) ([]certificates.Certificate, error) {
	all := inst.ListAll()
	for _, solution := range all {
		certs, err := solution.Load(raw)
		if err == nil && certs != nil {
			return certs, err
		}
	}
	return nil, fmt.Errorf("unsupported raw data")
}

// MakeChain ...
func (inst *CertSolutionManager) MakeChain(certs []certificates.Certificate) (certificates.Chain, error) {
	builder := certChainBuilder{}
	builder.init(inst.Roots)
	for _, cert := range certs {
		builder.add(cert)
	}
	return builder.make()
}

// LoadChain ...
func (inst *CertSolutionManager) LoadChain(raw *certificates.Raw) (certificates.Chain, error) {
	certs, err := inst.Load(raw)
	if err != nil {
		return nil, err
	}
	return inst.MakeChain(certs)
}

// Verify ...
func (inst *CertSolutionManager) Verify(chain certificates.Chain, op *certificates.VerifyOptions) error {
	if chain == nil {
		return fmt.Errorf("chain is nil")
	}
	solution := chain.GetCertificate().GetSolution()
	return solution.Verify(chain, op)
}

// GetInfo ...
func (inst *CertSolutionManager) GetInfo() *certificates.SolutionInfo {
	return &certificates.SolutionInfo{
		Algorithm:   "default",
		ContentType: "default",
		Format:      "default",
	}
}
