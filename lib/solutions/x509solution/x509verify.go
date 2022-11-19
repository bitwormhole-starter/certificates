package x509solution

import (
	"crypto/x509"
	"fmt"
	"time"

	"bitwormhole.com/starter/certificates"
	"github.com/bitwormhole/starter/vlog"
)

type x509veri struct {
	options  certificates.VerifyOptions
	chain    certificates.Chain
	roots    certificates.RootManager
	rawChain []*x509.Certificate
}

func (inst *x509veri) verify() error {

	steps := make([]func() error, 0)
	steps = append(steps, inst.prepare)
	steps = append(steps, inst.verifyStandard)

	for _, fn := range steps {
		err := fn()
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *x509veri) prepare() error {

	src := inst.chain
	dst := make([]*x509.Certificate, 0)

	for ; src != nil; src = src.GetParent() {
		cer, ok := src.GetCertificate().GetEntity().(*x509cert)
		if !ok {
			return fmt.Errorf("not a x509 cert")
		}
		dst = append(dst, cer.raw)
	}

	inst.rawChain = dst
	return nil
}

func (inst *x509veri) verifyStandard() error {

	chain := inst.rawChain
	if chain == nil {
		return fmt.Errorf("raw chain is nil")
	}

	size := len(chain)
	if size < 1 {
		return fmt.Errorf("empty cert chain")
	}
	targetCert := chain[0]
	mid := make([]*x509.Certificate, 0)
	for i, cer := range chain {
		if 0 < i {
			mid = append(mid, cer)
		}
	}

	opts := x509.VerifyOptions{}
	inst.prepareVerifiyOptionsWithIntermediates(&opts, mid)
	inst.prepareVerifiyOptionsWithRoots(&opts)
	opts.CurrentTime = time.Now()

	chs, err := targetCert.Verify(opts)
	if err != nil {
		return err
	}
	if chs == nil {
		return fmt.Errorf("chains is nil")
	}
	count := len(chs)
	vlog.Debug("chains.count=", count)
	return nil
}

func (inst *x509veri) prepareVerifiyOptionsWithRoots(opts *x509.VerifyOptions) {
	pool := x509.NewCertPool()
	src := inst.roots.ListAll()
	for _, root := range src {
		x509c, ok := root.GetEntity().(*x509cert)
		if ok && x509c != nil {
			raw := x509c.raw
			if raw != nil {
				pool.AddCert(raw)
			}
		}
	}
	opts.Roots = pool
}

func (inst *x509veri) prepareVerifiyOptionsWithIntermediates(opts *x509.VerifyOptions, src []*x509.Certificate) {
	pool := x509.NewCertPool()
	for _, cer := range src {
		pool.AddCert(cer)
	}
	opts.Intermediates = pool
}
