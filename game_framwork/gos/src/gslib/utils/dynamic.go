package utils

import "reflect"

func Call(instance interface{}, method string, params []reflect.Value) []reflect.Value {
	return reflect.ValueOf(instance).MethodByName(method).Call(params)
}

func CallWithArgs(instance interface{}, method string, params ...interface{}) []reflect.Value {
	in := ToReflectValues(params)
	return reflect.ValueOf(instance).MethodByName(method).Call(in)
}

func ToReflectValues(args []interface{}) []reflect.Value {
	in := make([]reflect.Value, len(args))
	for k, arg := range args {
		in[k] = reflect.ValueOf(arg)
	}
	return in
}
