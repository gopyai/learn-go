package main

import (
	"reflect"
	"encoding/json"
	"errors"
	"github.com/julienschmidt/httprouter"
	"fmt"
	"devx/iferr"
)

type (
	dataIn struct {
		Msg string
	}
	dataOut struct {
		Ret string
	}
)

func main() {
	APIServer(func(p httprouter.Params, in *dataIn) (out, ret interface{}) {
		fmt.Println("In:", in)
		return dataOut{"Asik"}, "Hello"
	})
}

// f signature:
// func(p httprouter.Params, in *struct) (out interface, ret interface)
func APIServer(f interface{}) {
	err := errors.New("api server error")

	vFunc := reflect.ValueOf(f)
	tFunc := vFunc.Type()
	if vFunc.Kind() != reflect.Func || tFunc.NumIn() != 2 || tFunc.NumOut() != 2 {
		panic(err)
	}

	// Verify function inputs
	tPar := tFunc.In(0)
	if tPar.String() != "httprouter.Params" {
		panic(err)
	}
	tIn := tFunc.In(1)
	if tIn.Kind() != reflect.Ptr || tIn.Elem().Kind() != reflect.Struct {
		panic(err)
	}

	// Verify function outputs
	if tFunc.Out(0).Kind() != reflect.Interface || tFunc.Out(1).Kind() != reflect.Interface {
		panic(err)
	}

	// Create new out to pass to function
	vIn := reflect.New(tIn.Elem())

	b := []byte(`{"Msg":"Hello"}`)
	iferr.Panic(json.Unmarshal(b, vIn.Interface()))

	// Call handler
	vOut := vFunc.Call([]reflect.Value{reflect.ValueOf(httprouter.Params{}), vIn})
	out := vOut[0].Interface()
	ret := vOut[1].Interface()

	b, err = json.Marshal(out)
	iferr.Panic(err)

	fmt.Println("Out:", string(b))
	fmt.Println("Ret:", ret)
}
