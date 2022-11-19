package certificates

// Chain ... 表示证书连上的一个节点
type Chain interface {
	GetParent() Chain
	GetCertificate() Certificate
}
