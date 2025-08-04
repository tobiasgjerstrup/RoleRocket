package main

/*
#cgo CFLAGS: -g -Wall
#include "mylib.h"
#include "mylib.c"
*/
import "C"
import "fmt"

func main() {
	C.say_hello()
	number := C.return_number()
	fmt.Println(number)

	// Proper capitalization and conversion
	fmt.Println(C.GoString(C.return_char()))

	number = C.return_another_number()
	fmt.Println(number)

	fmt.Println(C.GoString(C.return_char_again()))
}
