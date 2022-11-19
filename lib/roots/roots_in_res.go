package roots

import (
	"strings"

	"bitwormhole.com/starter/certificates"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/vlog"
)

// ResRoots ...
type ResRoots struct {
	markup.Component `class:"certificate-root-registry"`

	Context  application.Context `inject:"context"`
	Enabled  bool                `inject:"${certificates.roots-enabled}"`
	NameList string              `inject:"${certificates.roots}"`

	cached []*certificates.RootRegistration
}

func (inst *ResRoots) _Impl() certificates.RootRegistry {
	return inst
}

// ListRegistrations ...
func (inst *ResRoots) ListRegistrations() []*certificates.RootRegistration {
	src := inst.getAll()
	dst := make([]*certificates.RootRegistration, 0)
	for _, item1 := range src {
		item2 := inst.cloneItem(item1)
		dst = append(dst, item2)
	}
	return dst
}

func (inst *ResRoots) cloneItem(src *certificates.RootRegistration) *certificates.RootRegistration {
	dst := &certificates.RootRegistration{}
	if src != nil {
		dst.Title = src.Title
		dst.Enabled = src.Enabled
		raw := src.Raw
		if raw != nil {
			dst.Raw = raw.Clone()
		}
	}
	return dst
}

func (inst *ResRoots) getAll() []*certificates.RootRegistration {
	all := inst.cached
	if all == nil {
		all = inst.loadAll()
		inst.cached = all
	}
	return all
}

func (inst *ResRoots) loadAll() []*certificates.RootRegistration {
	dst := make([]*certificates.RootRegistration, 0)
	namelist := strings.Split(inst.NameList, ",")
	for _, name := range namelist {
		rr, err := inst.loadRoot(name)
		if err != nil {
			vlog.Warn(err)
		} else if rr != nil {
			dst = append(dst, rr)
		}
	}
	return dst
}

func (inst *ResRoots) loadRoot(name string) (*certificates.RootRegistration, error) {

	const prefixCert = "certificate."
	prefix := prefixCert + strings.TrimSpace(name)
	ctx := inst.Context
	props := ctx.GetProperties()
	getter := props.Getter()
	rr := &certificates.RootRegistration{}

	enabled := getter.GetBool(prefix+".enabled", true)
	src, err := props.GetPropertyRequired(prefix + ".src")
	if err != nil {
		return nil, err
	}

	bin, err := ctx.GetResources().GetBinary(src)
	if err != nil {
		return nil, err
	}

	rr.Enabled = enabled
	rr.Title = name
	rr.Raw = &certificates.Raw{Data: bin}
	return rr, nil
}
