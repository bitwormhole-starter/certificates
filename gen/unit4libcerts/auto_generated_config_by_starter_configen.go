// (todo:gen2.template) 
// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package unit4libcerts

import (
	certificates0xb7c330 "bitwormhole.com/starter/certificates"
	unit0x9449e7 "bitwormhole.com/starter/certificates/src/test/golang/unit"
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

	// component: com0-unit0x9449e7.TestPoint1
	cominfobuilder.Next()
	cominfobuilder.ID("com0-unit0x9449e7.TestPoint1").Class("life").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComTestPoint1{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}



    return nil
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4pComTestPoint1 : the factory of component: com0-unit0x9449e7.TestPoint1
type comFactory4pComTestPoint1 struct {

    mPrototype * unit0x9449e7.TestPoint1

	
	mContextSelector config.InjectionSelector
	mSolutionsSelector config.InjectionSelector
	mCertFileSelector config.InjectionSelector

}

func (inst * comFactory4pComTestPoint1) init() application.ComponentFactory {

	
	inst.mContextSelector = config.NewInjectionSelector("context",nil)
	inst.mSolutionsSelector = config.NewInjectionSelector("#certificate-solution-manager",nil)
	inst.mCertFileSelector = config.NewInjectionSelector("${test.cert.file}",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComTestPoint1) newObject() * unit0x9449e7.TestPoint1 {
	return & unit0x9449e7.TestPoint1 {}
}

func (inst * comFactory4pComTestPoint1) castObject(instance application.ComponentInstance) * unit0x9449e7.TestPoint1 {
	return instance.Get().(*unit0x9449e7.TestPoint1)
}

func (inst * comFactory4pComTestPoint1) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComTestPoint1) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComTestPoint1) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComTestPoint1) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComTestPoint1) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComTestPoint1) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Context = inst.getterForFieldContextSelector(context)
	obj.Solutions = inst.getterForFieldSolutionsSelector(context)
	obj.CertFile = inst.getterForFieldCertFileSelector(context)
	return context.LastError()
}

//getterForFieldContextSelector
func (inst * comFactory4pComTestPoint1) getterForFieldContextSelector (context application.InstanceContext) application.Context {
    return context.Context()
}

//getterForFieldSolutionsSelector
func (inst * comFactory4pComTestPoint1) getterForFieldSolutionsSelector (context application.InstanceContext) certificates0xb7c330.SolutionManager {

	o1 := inst.mSolutionsSelector.GetOne(context)
	o2, ok := o1.(certificates0xb7c330.SolutionManager)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com0-unit0x9449e7.TestPoint1")
		eb.Set("field", "Solutions")
		eb.Set("type1", "?")
		eb.Set("type2", "certificates0xb7c330.SolutionManager")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldCertFileSelector
func (inst * comFactory4pComTestPoint1) getterForFieldCertFileSelector (context application.InstanceContext) string {
    return inst.mCertFileSelector.GetString(context)
}




