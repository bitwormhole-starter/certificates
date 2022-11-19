package certificates

import "time"

// CertificateInfo ...
type CertificateInfo struct {
	Version           string
	SN                string
	StartedAt         time.Time
	StoppedAt         time.Time
	FingerprintSHA1   []byte
	FingerprintSHA256 []byte

	Subject Name
	Issuer  Name
}

// Certificate ... 表示一张证书
type Certificate interface {
	GetSolution() Solution
	GetInfo() *CertificateInfo
	GetEntity() CertificateEntity
	Equals(other Certificate) bool
}

// CertificateEntity ...
type CertificateEntity interface {
	Certificate() Certificate
}
