package tbs

import (
	"errors"
	"math/rand"
	"reflect"
	"time"
	//"fmt"
)

func MakeRand(max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max)
}

func Merge(ba []byte, bb []byte) []byte {
	buf := make([]byte, len(ba)+len(bb))
	copy(buf[:], ba)
	copy(buf[len(buf):], ba)

	return buf
}

func Call(m interface{}, name string, params ...interface{}) (result []reflect.Value, err error) {
	f := reflect.ValueOf(m)
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}

	method := f.MethodByName(name)
	if method.IsValid() {
		result = f.MethodByName(name).Call(in)
	} else {
		result = nil
		err = errors.New("method["+name+"] is not exist.")
	}

	return
}

func Get(m interface {}, name string) interface {}{
	f := reflect.ValueOf(m)
	filed := f.FieldByName(name)
	return filed.Interface()
}

//=====Injector=====
type Injector struct{
	valueMap  map[string]interface{}
}

func CreateInjector() Injector{
	instance := Injector{}
	instance.valueMap = make(map[string]interface {})

	return instance
}

func (this Injector) Set(key string, value interface{}) Injector {
	this.valueMap[key] = value
	return this
}

func (this Injector) Clear() Injector {
	for k := range this.valueMap {
		delete(this.valueMap, k)
	}
	return this
}

func (this Injector) Apply(target interface{}) {
	elem := reflect.ValueOf(target)
	for elem.Kind() == reflect.Ptr {
		elem = elem.Elem()
	}
	//fmt.Println(elem, elem.Kind())

	if elem.Kind() != reflect.Struct {
		return
	}

	//fmt.Println(elem, elem.NumField())
	t := elem.Type()
	fieldTagMap := make(map[reflect.StructTag]string)
	for i := 0; i < elem.NumField(); i++ {
		structField := t.Field(i)
		//fmt.Println(elem.Field(i).CanSet())
		if elem.Field(i).CanSet(){
			fieldTagMap[structField.Tag] = structField.Name
		}
	}

	for key, value := range this.valueMap{
		for k, v := range fieldTagMap{
			if k == reflect.StructTag(key) {
				elem.FieldByName(v).Set(reflect.ValueOf(value))
				break;
			}
		}
	}
}
