package main

//#include "sayhello.c"
import "C"
import "unsafe"

func main() {
	// Get a pointer to a []char.
	cs := C.CString("Hello from Go and C\n")

	// Say hello.
	C.sayhello(cs)

	// Strings are allocated with malloc so we need
	// to manually free that memory or face the consequences..
	C.free(unsafe.Pointer(cs))
}
