package main

/*
#include "sayhello.c"
*/
import "C"
import "unsafe"

func main() {
	cs := C.CString("Hello from Go and C\n")
	C.sayhello(cs)
	C.free(unsafe.Pointer(cs))
}
