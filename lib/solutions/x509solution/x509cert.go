package x509solution

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"crypto/x509/pkix"
	"fmt"
	"hash"
	"strconv"
	"strings"

	"bitwormhole.com/starter/certificates"
)

type x509cert struct {
	solution certificates.Solution
	raw      *x509.Certificate
	info     certificates.CertificateInfo
	title    string
}

func (inst *x509cert) _Impl() (certificates.Certificate, certificates.CertificateEntity) {
	return inst, inst
}

func (inst *x509cert) init(c1 *x509.Certificate) error {
	if c1 == nil {
		return fmt.Errorf("param: [x509.Certificate] is nil")
	}
	inst.loadInfo(c1)
	inst.raw = c1
	inst.title = inst.info.Subject.String()
	return nil
}

func (inst *x509cert) loadInfo(c1 *x509.Certificate) {
	t1 := c1.NotBefore
	t2 := c1.NotAfter
	issuer := inst.convertName(c1.Issuer)
	subject := inst.convertName(c1.Subject)
	version := strconv.Itoa(c1.Version)
	sn := c1.SerialNumber.String()
	fp1 := inst.computeFingerprint(c1, sha1.New())
	fp256 := inst.computeFingerprint(c1, sha256.New())
	inst.info = certificates.CertificateInfo{
		StartedAt:         t1,
		StoppedAt:         t2,
		Issuer:            issuer,
		Subject:           subject,
		SN:                sn,
		Version:           version,
		FingerprintSHA1:   fp1,
		FingerprintSHA256: fp256,
	}
}

func (inst *x509cert) computeFingerprint(c1 *x509.Certificate, h hash.Hash) []byte {
	data := c1.Raw
	h.Reset()
	h.Write(data)
	return h.Sum(nil)
}

func (inst *x509cert) Certificate() certificates.Certificate {
	return inst
}

func (inst *x509cert) GetSolution() certificates.Solution {
	return inst.solution
}

func (inst *x509cert) GetInfo() *certificates.CertificateInfo {
	dst := &certificates.CertificateInfo{}
	*dst = inst.info
	return dst
}

func (inst *x509cert) GetEntity() certificates.CertificateEntity {
	return inst
}

func (inst *x509cert) Equals(other certificates.Certificate) bool {
	if other == nil {
		return false
	}
	cer2, ok := other.GetEntity().(*x509cert)
	if !ok {
		return false
	}

	cer2.GetInfo() // todo ...
	return false
}

func (inst *x509cert) convertName(src pkix.Name) certificates.Name {

	// Country, Organization, OrganizationalUnit []string
	// Locality, Province                        []string
	// StreetAddress, PostalCode                 []string
	// SerialNumber, CommonName                  string

	const sep = ","
	dst := certificates.Name{}

	dst.Country = strings.Join(src.Country, sep)
	dst.Organization = strings.Join(src.Organization, sep)
	dst.OrganizationalUnit = strings.Join(src.OrganizationalUnit, sep)

	dst.Locality = strings.Join(src.Locality, sep)
	dst.Province = strings.Join(src.Province, sep)

	dst.StreetAddress = strings.Join(src.StreetAddress, sep)
	dst.PostalCode = strings.Join(src.PostalCode, sep)

	dst.CommonName = src.CommonName
	dst.SerialNumber = src.SerialNumber

	return dst
}
