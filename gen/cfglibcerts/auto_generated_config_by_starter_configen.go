// (todo:gen2.template) 
// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package cfglibcerts

import (
	certificates0xb7c330 "bitwormhole.com/starter/certificates"
	lib0x47d8cf "bitwormhole.com/starter/certificates/lib"
	roots0x5e4213 "bitwormhole.com/starter/certificates/lib/roots"
	x509solution0x7650e0 "bitwormhole.com/starter/certificates/lib/solutions/x509solution"
	application "github.com/bitwormhole/starter/application"
	config "github.com/bitwormhole/starter/application/config"
	lang "github.com/bitwormhole/starter/lang"
	util "github.com/bitwormhole/starter/util"
    
)


func nop(x ... interface{}){
	util.Int64ToTime(0)
	lang.CreateReleasePool()
}


func autoGenConfig(cb application.ConfigBuilder) error {

	var err error = nil
	cominfobuilder := config.ComInfo()
	nop(err,cominfobuilder)

	// component: com0-roots0x5e4213.ResRoots
	cominfobuilder.Next()
	cominfobuilder.ID("com0-roots0x5e4213.ResRoots").Class("certificate-root-registry").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComResRoots{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: certificate-root-manager
	cominfobuilder.Next()
	cominfobuilder.ID("certificate-root-manager").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComRootCertManager{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com2-x509solution0x7650e0.RsaX509Solution
	cominfobuilder.Next()
	cominfobuilder.ID("com2-x509solution0x7650e0.RsaX509Solution").Class("certificate-solution-registry").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComRsaX509Solution{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: certificate-solution-manager
	cominfobuilder.Next()
	cominfobuilder.ID("certificate-solution-manager").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComCertSolutionManager{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}



    return nil
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4pComResRoots : the factory of component: com0-roots0x5e4213.ResRoots
type comFactory4pComResRoots struct {

    mPrototype * roots0x5e4213.ResRoots

	
	mContextSelector config.InjectionSelector
	mEnabledSelector config.InjectionSelector
	mNameListSelector config.InjectionSelector

}

func (inst * comFactory4pComResRoots) init() application.ComponentFactory {

	
	inst.mContextSelector = config.NewInjectionSelector("context",nil)
	inst.mEnabledSelector = config.NewInjectionSelector("${certificates.roots-enabled}",nil)
	inst.mNameListSelector = config.NewInjectionSelector("${certificates.roots}",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComResRoots) newObject() * roots0x5e4213.ResRoots {
	return & roots0x5e4213.ResRoots {}
}

func (inst * comFactory4pComResRoots) castObject(instance application.ComponentInstance) * roots0x5e4213.ResRoots {
	return instance.Get().(*roots0x5e4213.ResRoots)
}

func (inst * comFactory4pComResRoots) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComResRoots) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComResRoots) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComResRoots) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComResRoots) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComResRoots) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Context = inst.getterForFieldContextSelector(context)
	obj.Enabled = inst.getterForFieldEnabledSelector(context)
	obj.NameList = inst.getterForFieldNameListSelector(context)
	return context.LastError()
}

//getterForFieldContextSelector
func (inst * comFactory4pComResRoots) getterForFieldContextSelector (context application.InstanceContext) application.Context {
    return context.Context()
}

//getterForFieldEnabledSelector
func (inst * comFactory4pComResRoots) getterForFieldEnabledSelector (context application.InstanceContext) bool {
    return inst.mEnabledSelector.GetBool(context)
}

//getterForFieldNameListSelector
func (inst * comFactory4pComResRoots) getterForFieldNameListSelector (context application.InstanceContext) string {
    return inst.mNameListSelector.GetString(context)
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComRootCertManager : the factory of component: certificate-root-manager
type comFactory4pComRootCertManager struct {

    mPrototype * lib0x47d8cf.RootCertManager

	
	mSolutionsSelector config.InjectionSelector
	mRegsSelector config.InjectionSelector

}

func (inst * comFactory4pComRootCertManager) init() application.ComponentFactory {

	
	inst.mSolutionsSelector = config.NewInjectionSelector("#certificate-solution-manager",nil)
	inst.mRegsSelector = config.NewInjectionSelector(".certificate-root-registry",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComRootCertManager) newObject() * lib0x47d8cf.RootCertManager {
	return & lib0x47d8cf.RootCertManager {}
}

func (inst * comFactory4pComRootCertManager) castObject(instance application.ComponentInstance) * lib0x47d8cf.RootCertManager {
	return instance.Get().(*lib0x47d8cf.RootCertManager)
}

func (inst * comFactory4pComRootCertManager) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComRootCertManager) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComRootCertManager) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComRootCertManager) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRootCertManager) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRootCertManager) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Solutions = inst.getterForFieldSolutionsSelector(context)
	obj.Regs = inst.getterForFieldRegsSelector(context)
	return context.LastError()
}

//getterForFieldSolutionsSelector
func (inst * comFactory4pComRootCertManager) getterForFieldSolutionsSelector (context application.InstanceContext) certificates0xb7c330.SolutionManager {

	o1 := inst.mSolutionsSelector.GetOne(context)
	o2, ok := o1.(certificates0xb7c330.SolutionManager)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "certificate-root-manager")
		eb.Set("field", "Solutions")
		eb.Set("type1", "?")
		eb.Set("type2", "certificates0xb7c330.SolutionManager")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldRegsSelector
func (inst * comFactory4pComRootCertManager) getterForFieldRegsSelector (context application.InstanceContext) []certificates0xb7c330.RootRegistry {
	list1 := inst.mRegsSelector.GetList(context)
	list2 := make([]certificates0xb7c330.RootRegistry, 0, len(list1))
	for _, item1 := range list1 {
		item2, ok := item1.(certificates0xb7c330.RootRegistry)
		if ok {
			list2 = append(list2, item2)
		}
	}
	return list2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComRsaX509Solution : the factory of component: com2-x509solution0x7650e0.RsaX509Solution
type comFactory4pComRsaX509Solution struct {

    mPrototype * x509solution0x7650e0.RsaX509Solution

	
	mRootsSelector config.InjectionSelector

}

func (inst * comFactory4pComRsaX509Solution) init() application.ComponentFactory {

	
	inst.mRootsSelector = config.NewInjectionSelector("#certificate-root-manager",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComRsaX509Solution) newObject() * x509solution0x7650e0.RsaX509Solution {
	return & x509solution0x7650e0.RsaX509Solution {}
}

func (inst * comFactory4pComRsaX509Solution) castObject(instance application.ComponentInstance) * x509solution0x7650e0.RsaX509Solution {
	return instance.Get().(*x509solution0x7650e0.RsaX509Solution)
}

func (inst * comFactory4pComRsaX509Solution) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComRsaX509Solution) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComRsaX509Solution) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComRsaX509Solution) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRsaX509Solution) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRsaX509Solution) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Roots = inst.getterForFieldRootsSelector(context)
	return context.LastError()
}

//getterForFieldRootsSelector
func (inst * comFactory4pComRsaX509Solution) getterForFieldRootsSelector (context application.InstanceContext) certificates0xb7c330.RootManager {

	o1 := inst.mRootsSelector.GetOne(context)
	o2, ok := o1.(certificates0xb7c330.RootManager)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com2-x509solution0x7650e0.RsaX509Solution")
		eb.Set("field", "Roots")
		eb.Set("type1", "?")
		eb.Set("type2", "certificates0xb7c330.RootManager")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComCertSolutionManager : the factory of component: certificate-solution-manager
type comFactory4pComCertSolutionManager struct {

    mPrototype * lib0x47d8cf.CertSolutionManager

	
	mRootsSelector config.InjectionSelector
	mRegsSelector config.InjectionSelector

}

func (inst * comFactory4pComCertSolutionManager) init() application.ComponentFactory {

	
	inst.mRootsSelector = config.NewInjectionSelector("#certificate-root-manager",nil)
	inst.mRegsSelector = config.NewInjectionSelector(".certificate-solution-registry",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComCertSolutionManager) newObject() * lib0x47d8cf.CertSolutionManager {
	return & lib0x47d8cf.CertSolutionManager {}
}

func (inst * comFactory4pComCertSolutionManager) castObject(instance application.ComponentInstance) * lib0x47d8cf.CertSolutionManager {
	return instance.Get().(*lib0x47d8cf.CertSolutionManager)
}

func (inst * comFactory4pComCertSolutionManager) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComCertSolutionManager) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComCertSolutionManager) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComCertSolutionManager) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComCertSolutionManager) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComCertSolutionManager) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Roots = inst.getterForFieldRootsSelector(context)
	obj.Regs = inst.getterForFieldRegsSelector(context)
	return context.LastError()
}

//getterForFieldRootsSelector
func (inst * comFactory4pComCertSolutionManager) getterForFieldRootsSelector (context application.InstanceContext) certificates0xb7c330.RootManager {

	o1 := inst.mRootsSelector.GetOne(context)
	o2, ok := o1.(certificates0xb7c330.RootManager)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "certificate-solution-manager")
		eb.Set("field", "Roots")
		eb.Set("type1", "?")
		eb.Set("type2", "certificates0xb7c330.RootManager")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldRegsSelector
func (inst * comFactory4pComCertSolutionManager) getterForFieldRegsSelector (context application.InstanceContext) []certificates0xb7c330.SolutionRegistry {
	list1 := inst.mRegsSelector.GetList(context)
	list2 := make([]certificates0xb7c330.SolutionRegistry, 0, len(list1))
	for _, item1 := range list1 {
		item2, ok := item1.(certificates0xb7c330.SolutionRegistry)
		if ok {
			list2 = append(list2, item2)
		}
	}
	return list2
}




