package main

/*
#cgo CFLAGS: -Imruby/include
#cgo LDFLAGS: -Lmruby/build/host/lib -lmruby
#include <stdlib.h>
#include "mruby.h"
#include "mruby/compile.h"
*/
import "C"

import (
	"flag"
	"io/ioutil"
	"os"
	"unsafe"
)

func main() {
	flag.Parse()
	args := flag.Args()

	var err error
	var fp *os.File
	if len(args) == 0 {
		fp = os.Stdin
	} else {
		fp, err = os.Open(args[0])
		if err != nil {
			panic(err)
		}
		defer fp.Close()
	}
	code, err := ioutil.ReadAll(fp)
	if err != nil {
		panic(err)
	}

	mrb := C.mrb_open()
	if mrb == nil {
		panic("Can not open mruby.")
	}

	ccode := C.CString(string(code))
	C.mrb_load_string(mrb, ccode)
	C.mrb_close(mrb)
	C.free(unsafe.Pointer(ccode))
}
