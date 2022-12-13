package main

import (
	"fmt"
	"reflect"
)

var pl = fmt.Println

func main() {
	pl(reflect.TypeOf(25))
	pl(reflect.TypeOf("asfda"))
	pl(reflect.TypeOf(true))
	pl(reflect.TypeOf("Hello"))
	pl(reflect.TypeOf(("*")))
}
