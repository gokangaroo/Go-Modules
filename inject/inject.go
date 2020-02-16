package main

import (
	"fmt"
	"reflect"
)

type Injector interface {
	Applicator //注册一个结构体
	Invoker    //实际调用
	TypeMapper //类型容器

	SetParent(Injector)
}

//将依赖注入的容器放到一个结构体中 根据这个结构体是否有'inject'标签
type Applicator interface {
	Apply(interface{}) error
}

//具体依赖注入调用的方法
type Invoker interface {
	Invoke(interface{}) ([]reflect.Value, error)
}

//用来作为依赖注入容器的 方法都是链式调用
type TypeMapper interface {
	//设置一个对象
	Map(interface{}) TypeMapper
	//将一个对象注入到一个接口中 type:接口 value:对象
	MapTo(interface{}, interface{}) TypeMapper
	//手动设置key和value
	Set(reflect.Type, reflect.Value) TypeMapper
	//从容器中的获取某个类型的注入对象
	Get(reflect.Type) reflect.Value
}

//实际的的容器 实现Injector所有方法
type injector struct {
	//容器map
	values map[reflect.Type]reflect.Value

	//设置一个parent 即是可以嵌套的
	parent Injector
}

// InterfaceOf dereferences a pointer to an Interface type.
func InterfaceOf(value interface{}) reflect.Type {
	t := reflect.TypeOf(value)
	//如果是指针类型
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() != reflect.Interface {
		panic("Called inject.InterfaceOf with a value that is not a pointer to an interface. (*MyInterface)(nil)")
	}
	return t
}

func New() Injector {
	return &injector{
		values: make(map[reflect.Type]reflect.Value),
	}
}

func (inject *injector) Invoke(f interface{}) ([]reflect.Value, error) {
	t := reflect.TypeOf(f)

	var in = make([]reflect.Value, t.NumIn()) //Panic if t is not kind of Func
	for i := 0; i < t.NumIn(); i++ {
		argType := t.In(i)
		val := inject.Get(argType)
		if !val.IsValid() {
			return nil, fmt.Errorf("Value not found for type %v", argType)
		}

		in[i] = val
	}

	return reflect.ValueOf(f).Call(in), nil
}

// Maps dependencies in the Type map to each field in the struct
// that is tagged with 'inject'.
// Returns an error if the injection fails.
func (injector *injector) Apply(val interface{}) error {
	v := reflect.ValueOf(val)

	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil // Should not panic here ?
	}

	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		structField := t.Field(i)
		if f.CanSet() && (structField.Tag == "inject" || structField.Tag.Get("inject") != "") {
			ft := f.Type()
			v := injector.Get(ft)
			if !v.IsValid() {
				return fmt.Errorf("Value not found for type %v", ft)
			}

			f.Set(v)
		}

	}

	return nil
}

// Maps the concrete value of val to its dynamic type using reflect.TypeOf,
// It returns the TypeMapper registered in.
func (i *injector) Map(val interface{}) TypeMapper {
	i.values[reflect.TypeOf(val)] = reflect.ValueOf(val)
	return i
}

func (i *injector) MapTo(val interface{}, ifacePtr interface{}) TypeMapper {
	i.values[InterfaceOf(ifacePtr)] = reflect.ValueOf(val)
	return i
}

// Maps the given reflect.Type to the given reflect.Value and returns
// the Typemapper the mapping has been registered in.
func (i *injector) Set(typ reflect.Type, val reflect.Value) TypeMapper {
	i.values[typ] = val
	return i
}

func (i *injector) Get(t reflect.Type) reflect.Value {
	val := i.values[t]

	if val.IsValid() {
		return val
	}

	// no concrete types found, try to find implementors
	// if t is an interface
	if t.Kind() == reflect.Interface {
		for k, v := range i.values {
			if k.Implements(t) {
				val = v
				break
			}
		}
	}

	// Still no type found, try to look it up on the parent
	if !val.IsValid() && i.parent != nil {
		val = i.parent.Get(t)
	}

	return val

}

func (i *injector) SetParent(parent Injector) {
	i.parent = parent
}
