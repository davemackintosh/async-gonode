package main

/*
#include <stdio.h>
#include <stdlib.h>

void myprint(char* s) {
  printf("%s", s);
}
*/
import "C"
import "unsafe"

func main() {
	cs := C.CString("Hello from Go and C\n")
	C.myprint(cs)
	C.free(unsafe.Pointer(cs))
}
