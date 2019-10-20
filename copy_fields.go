package gotoo

import (
	"fmt"
	"reflect"
)

func CopyFields(source interface{}, target interface{}) (interface{}, error) {
	sv, tv := reflect.ValueOf(source), reflect.ValueOf(target)
	o := reflect.New(reflect.TypeOf(target))
	if sv.Type().Kind() != reflect.Struct || tv.Type().Kind() != reflect.Struct {
		return nil, fmt.Errorf("Invalid Type Error: Input type should be struct")
	}
	for i := 0; i < reflect.TypeOf(target).NumField(); i++ {
		sfv := sv.FieldByName(tv.Type().Field(i).Name)
		o.Elem().Field(i).Set(tv.Field(i))
		if sfv.IsValid() && tv.Field(i).Type() == sfv.Type() {
			o.Elem().Field(i).Set(sfv)
		}
	}
	return o.Interface(), nil
}

