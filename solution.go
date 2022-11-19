package certificates

// SolutionInfo 方案信息
type SolutionInfo struct {
	Algorithm   string // like 'RSA'
	ContentType string // like 'block'
	Format      string // like 'x509'
}

// VerifyOptions ...
type VerifyOptions struct {
	IgnoreTime    bool
	IgnoreRoot    bool
	IgnoreParents bool
}

// Solution ... 表示一套证书方案
type Solution interface {
	Load(raw *Raw) ([]Certificate, error)
	Verify(chain Chain, op *VerifyOptions) error
	GetInfo() *SolutionInfo
}

// SolutionRegistration ...证书方案注册信息
type SolutionRegistration struct {
	Info     SolutionInfo
	Solution Solution
}

// SolutionRegistry ... 证书方案注册人
// [inject:".certificate-solution-registry"]
type SolutionRegistry interface {
	GetRegistration() *SolutionRegistration
}

// SolutionManager 表示证书方案管理器
// [inject:"#certificate-solution-manager"]
type SolutionManager interface {
	Solution

	MakeChain(certs []Certificate) (Chain, error)

	LoadChain(raw *Raw) (Chain, error)

	ListAll() []Solution
}
