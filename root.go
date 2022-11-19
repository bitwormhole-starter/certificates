package certificates

// Root 表示根证书
type Root interface {
	Certificate
	Title() string
}

// RootRegistration 表示根证书注册信息
type RootRegistration struct {
	Title   string
	Raw     *Raw
	Enabled bool
}

// RootRegistry 表示根证书注册对象
// [inject:".certificate-root-registry"]
type RootRegistry interface {
	ListRegistrations() []*RootRegistration
}

// RootManager 表示根证书管理器
// [inject:"#certificate-root-manager"]
type RootManager interface {
	ListAll() []Root
	Find(target *Name) (Root, error)
	Contains(cert Certificate) bool
}
