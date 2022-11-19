package lib

import (
	"fmt"
	"sort"

	"bitwormhole.com/starter/certificates"
)

type certChainBuilder struct {
	roots certificates.RootManager
	table map[string]*certHolder
}

func (inst *certChainBuilder) init(roots certificates.RootManager) {
	inst.roots = roots
	inst.table = make(map[string]*certHolder)
}

func (inst *certChainBuilder) getHolder(key string, create bool) *certHolder {
	h := inst.table[key]
	if h == nil && create {
		h = &certHolder{key: key}
		inst.table[key] = h
	}
	return h
}

func (inst *certChainBuilder) add(cert certificates.Certificate) {
	info := cert.GetInfo()
	key1 := info.Subject.String()    // child
	key2 := info.Issuer.String()     // parent
	h1 := inst.getHolder(key1, true) // child
	h2 := inst.getHolder(key2, true) // parent
	h1.cert = cert
	h1.subject = info.Subject
	h2.subject = info.Issuer
}

func (inst *certChainBuilder) make() (certificates.Chain, error) {

	err := inst.doFindRoot()
	if err != nil {
		return nil, err
	}

	err = inst.doMark()
	if err != nil {
		return nil, err
	}

	list := inst.doSort()
	return inst.makeChainWithList(list)
}

func (inst *certChainBuilder) makeChainWithList(list []*certHolder) (certificates.Chain, error) {
	src := list
	var p certificates.Chain = nil
	for _, holder := range src {
		cert := holder.cert
		if cert == nil {
			return nil, fmt.Errorf("cert for chain is nil")
		}
		next := &chainImpl{
			parent: p,
			cert:   cert,
		}
		p = next
	}
	if p == nil {
		return nil, fmt.Errorf("chain is nil")
	}
	return p, nil
}

// 查找 root
func (inst *certChainBuilder) doFindRoot() error {
	src := inst.table
	for _, holder := range src {
		if holder.cert == nil {
			target := holder.subject
			root, err := inst.roots.Find(&target)
			if err != nil {
				return err
			}
			holder.cert = root
		}
	}
	return nil
}

// 标记 refcnt
func (inst *certChainBuilder) doMark() error {

	src := inst.table
	dst := make([]*certHolder, 0)

	// reset all
	for _, holder := range src {
		holder.refcnt = 0
		dst = append(dst, holder)
	}

	// scan all
	for _, holder := range dst {
		err := inst.markChainWithHolder(holder)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *certChainBuilder) markChainWithHolder(h *certHolder) error {
	table := inst.table
	for h != nil {
		h.refcnt++
		cer := h.cert
		if cer == nil {
			break
		}
		info := cer.GetInfo()
		key2 := info.Issuer.String()
		if h.key == key2 {
			break
		}
		h = table[key2]
	}
	return nil
}

// 排序
func (inst *certChainBuilder) doSort() []*certHolder {
	src := inst.table
	list := &certHolderList{}
	for _, item := range src {
		list.append(item)
	}
	list.sort()
	return list.items
}

////////////////////////////////////////////////////////////////////////////////

type certHolder struct {
	key     string
	cert    certificates.Certificate
	subject certificates.Name
	refcnt  int
}

////////////////////////////////////////////////////////////////////////////////

type certHolderList struct {
	items []*certHolder
}

func (inst *certHolderList) append(h *certHolder) {
	inst.items = append(inst.items, h)
}

func (inst *certHolderList) sort() {
	sort.Sort(inst)
}

func (inst *certHolderList) Swap(a, b int) {
	ao := inst.items[a]
	bo := inst.items[b]
	inst.items[a] = bo
	inst.items[b] = ao
}

func (inst *certHolderList) Less(a, b int) bool {
	ao := inst.items[a]
	bo := inst.items[b]
	return ao.refcnt > bo.refcnt
}

func (inst *certHolderList) Len() int {
	return len(inst.items)
}

////////////////////////////////////////////////////////////////////////////////
