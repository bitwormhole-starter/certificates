package unit

import (
	"bitwormhole.com/starter/certificates"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
)

// TestPoint1 ...
type TestPoint1 struct {
	markup.Component `class:"life"`

	Context   application.Context          `inject:"context"`
	Solutions certificates.SolutionManager `inject:"#certificate-solution-manager"`
	CertFile  string                       `inject:"${test.cert.file}"`
}

func (inst *TestPoint1) _Impl() application.LifeRegistry {
	return inst
}

// GetLifeRegistration ...
func (inst *TestPoint1) GetLifeRegistration() *application.LifeRegistration {
	return &application.LifeRegistration{
		OnStart: inst.doTest,
	}
}

func (inst *TestPoint1) doTest() error {

	ctx := inst.Context
	path := inst.CertFile
	bin, err := ctx.GetResources().GetBinary(path)
	if err != nil {
		return err
	}

	raw := &certificates.Raw{Data: bin}
	chain, err := inst.Solutions.LoadChain(raw)
	if err != nil {
		return err
	}

	solution := chain.GetCertificate().GetSolution()
	err = solution.Verify(chain, &certificates.VerifyOptions{})
	if err != nil {
		return err
	}

	return nil
}
