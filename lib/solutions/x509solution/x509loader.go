package x509solution

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"

	"bitwormhole.com/starter/certificates"
)

type myRsaX509Loader struct {
	solution certificates.Solution
}

func (inst *myRsaX509Loader) load(data []byte) ([]certificates.Certificate, error) {
	blocks, err := inst.loadPemBlocks(data)
	if err != nil {
		return nil, err
	}
	return inst.loadCerts(blocks)
}

func (inst *myRsaX509Loader) loadCerts(src []*pem.Block) ([]certificates.Certificate, error) {
	dst := make([]certificates.Certificate, 0)
	for _, b := range src {
		cer, err := inst.loadCert(b)
		if err != nil {
			return nil, err
		}
		dst = append(dst, cer)
	}
	return dst, nil
}

func (inst *myRsaX509Loader) loadCert(b *pem.Block) (certificates.Certificate, error) {
	cer1, err := x509.ParseCertificate(b.Bytes)
	if err != nil {
		return nil, err
	}
	cer2 := &x509cert{solution: inst.solution}
	err = cer2.init(cer1)
	if err != nil {
		return nil, err
	}
	return cer2, nil
}

func (inst *myRsaX509Loader) loadPemBlocks(data []byte) ([]*pem.Block, error) {
	if data == nil {
		return nil, fmt.Errorf("data is nil")
	}
	blocks := make([]*pem.Block, 0)
	data2 := data
	for {
		// size := len(data2)
		// vlog.Info("size=", size)
		block, rest := pem.Decode(data2)
		if block != nil {
			blocks = append(blocks, block)
		} else {
			break
		}
		data2 = rest
	}
	if len(blocks) < 1 {
		return nil, fmt.Errorf("no PEM block")
	}
	return blocks, nil
}
