package lib

import (
	"fmt"

	"bitwormhole.com/starter/certificates"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/vlog"
)

// RootCertManager ...
type RootCertManager struct {
	markup.Component `id:"certificate-root-manager"`

	Solutions certificates.SolutionManager `inject:"#certificate-solution-manager"`
	Regs      []certificates.RootRegistry  `inject:".certificate-root-registry"`

	all []certificates.Root
}

func (inst *RootCertManager) _Impl() certificates.RootManager {
	return inst
}

// ListAll ...
func (inst *RootCertManager) ListAll() []certificates.Root {
	src := inst.all
	if src == nil {
		src = inst.loadAll()
		inst.all = src
	}
	dst := make([]certificates.Root, 0)
	dst = append(dst, src...)
	return dst
}

func (inst *RootCertManager) loadAll() []certificates.Root {
	src := inst.Regs
	dst := make([]certificates.Root, 0)
	for _, rr1 := range src {
		rrlist := rr1.ListRegistrations()
		for _, rr2 := range rrlist {
			root, err := inst.loadCert(rr2)
			if err == nil {
				if root != nil {
					dst = append(dst, root)
				}
			} else {
				vlog.Warn(err)
			}
		}
	}
	return dst
}

func (inst *RootCertManager) loadCert(rr *certificates.RootRegistration) (certificates.Root, error) {
	if rr == nil {
		return nil, nil
	}
	if !rr.Enabled {
		return nil, nil
	}
	raw := rr.Raw
	certs, err := inst.Solutions.Load(raw)
	if err != nil {
		return nil, err
	}
	if len(certs) != 1 {
		return nil, fmt.Errorf("bad root cert: %v", rr.Title)
	}
	cert := certs[0]
	root := &rootCert{
		cert:  cert,
		title: rr.Title,
	}
	return root, nil
}

// Find ...
func (inst *RootCertManager) Find(target *certificates.Name) (certificates.Root, error) {
	u1 := target
	all := inst.ListAll()
	for _, root := range all {
		certInfo := root.GetInfo()
		u2 := &certInfo.Subject
		if u1.Equals(u2) {
			return root, nil
		}
	}
	return nil, fmt.Errorf("no root cert with name:%v", u1.CommonName)
}

// Contains ...
func (inst *RootCertManager) Contains(cert certificates.Certificate) bool {
	all := inst.ListAll()
	for _, root := range all {
		if root.Equals(cert) {
			return true
		}
	}
	return false
}

////////////////////////////////////////////////////////////////////////////////

type rootCert struct {
	cert  certificates.Certificate
	title string
}

func (inst *rootCert) _Impl() certificates.Root {
	return inst
}

func (inst *rootCert) Title() string {
	return inst.title
}

func (inst *rootCert) GetEntity() certificates.CertificateEntity {
	return inst.cert.GetEntity()
}

func (inst *rootCert) GetSolution() certificates.Solution {
	return inst.cert.GetSolution()
}

func (inst *rootCert) GetInfo() *certificates.CertificateInfo {
	return inst.cert.GetInfo()
}

func (inst *rootCert) Equals(other certificates.Certificate) bool {
	return inst.cert.Equals(other)
}
