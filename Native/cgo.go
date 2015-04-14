package main

/*
#include "sayhello.c"
*/
import "C"
import "unsafe"

func main() {
	// Get a pointer to a []char*
	cs := C.CString("Hello from Go and C\n")

	// Say hello.
	C.sayhello(cs)

	// Use that pointer to free memory, this is C..
	C.free(unsafe.Pointer(cs))
}
