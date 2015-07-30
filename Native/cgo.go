package main

//#include "sayhello.c"
import "C"
import "unsafe"

func main() {
	// Get a pointer to a []char to pass to C.
	cs := C.CString("I'm sorry, world.")

	// Strings are allocated with malloc so we need
	// to manually free that memory or face the consequences
	// when we're done with it.
	defer C.free(unsafe.Pointer(cs))

	// Say hello.
	C.sayhello(cs)
}
